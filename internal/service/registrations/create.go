package registrations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/upikoth/starter-go/internal/models"
)

var RegistrationEmailTemplate = `
<table width="100%%" border="0" cellspacing="20" cellpadding="0"
	style="background: #fff; max-width: 600px; margin: auto; border-radius: 10px">
	<tbody>
		<tr>
			<td align="center" style="
					padding: 10px 0px;
					font-size: 22px;
					font-family: Helvetica, Arial, sans-serif;
					color: #444;
				">
				Регистрация на <strong>%s</strong>
			</td>
		</tr>
		<tr>
			<td align="center" style="padding: 20px 0">
				<table border="0" cellspacing="0" cellpadding="0">
					<tbody>
						<tr>
							<td align="center" style="border-radius: 3px" bgcolor="#1976D2">
								<a href="%s?token=%s" target="_blank" style="
										font-size: 18px;
										font-family: Helvetica, Arial, sans-serif;
										color: #fff;
										text-decoration: none;
										border-radius: 3px;
										padding: 10px 20px;
										border: 1px solid #1976D2;
										display: inline-block;
										font-weight: bold;
									" rel="noopener noreferrer">
									Подтвердить регистрацию
								</a>
							</td>
						</tr>
					</tbody>
				</table>
			</td>
		</tr>
		<tr>
			<td align="center" style="
					padding: 0px 0px 10px 0px;
					font-size: 16px;
					line-height: 22px;
					font-family: Helvetica, Arial, sans-serif;
					color: #444;
				">
				Если вы не отправляли заявку на регистрацию, игнорируйте сообщение.
			</td>
		</tr>
	</tbody>
</table>
`

func (r *Registrations) Create(
	inputCtx context.Context,
	params models.RegistrationCreateParams,
) (models.Registration, error) {
	span := sentry.StartSpan(inputCtx, "Service: Registrations.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	registration := models.Registration{
		ID:                uuid.New().String(),
		Email:             params.Email,
		ConfirmationToken: uuid.New().String(),
	}

	existingUser, err := r.repository.Ydb.Users.GetByEmail(ctx, registration.Email)

	if err != nil {
		sentry.CaptureException(err)
		return registration, &models.Error{
			Code:        models.ErrorCodeRegistrationYdbFindUser,
			Description: err.Error(),
		}
	}

	if existingUser.ID != "" {
		return registration, &models.Error{
			Code:        models.ErrorCodeRegistrationUserWithThisEmailAlreadyExist,
			Description: "A user with the specified email already exists",
			StatusCode:  http.StatusBadRequest,
		}
	}

	existingRegistration, err := r.repository.Ydb.Registrations.GetByEmail(ctx, registration.Email)

	if err != nil {
		sentry.CaptureException(err)
		return registration, &models.Error{
			Code:        models.ErrorCodeRegistrationYdbCreateRegistration,
			Description: err.Error(),
		}
	}

	if existingRegistration.ID != "" {
		registration = existingRegistration
	} else {
		registration, err = r.repository.Ydb.Registrations.Create(ctx, registration)
	}

	if err != nil {
		sentry.CaptureException(err)
		return registration, &models.Error{
			Code:        models.ErrorCodeRegistrationYdbCreateRegistration,
			Description: err.Error(),
		}
	}

	registrationEmail := fmt.Sprintf(
		RegistrationEmailTemplate,
		r.config.FrontURL,
		r.config.FrontConfirmationRegistrationURL,
		registration.ConfirmationToken,
	)

	err = r.repository.Ycp.SendEmail(
		ctx,
		registration.Email,
		"Регистрация на "+r.config.FrontURL,
		registrationEmail,
	)

	if err != nil {
		sentry.CaptureException(err)
		return registration, &models.Error{
			Code:        models.ErrorCodeRegistrationSMTPSendEmail,
			Description: err.Error(),
		}
	}

	return registration, nil
}
