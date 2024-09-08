package passwordrecoveryrequests

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/upikoth/starter-go/internal/models"
)

//nolint:gosec // тут нет хардкода паролей.
var PasswordRecoveryRequestEmailTemplate = `
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
				Восстановление пароля на <strong>%s</strong>
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
									Восстановить пароль
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
				Если вы не отправляли заявку на восстановление пароля, игнорируйте сообщение.
			</td>
		</tr>
	</tbody>
</table>
`

func (p *PasswordRecoveryRequests) Create(
	inputCtx context.Context,
	params models.PasswordRecoveryRequestCreateParams,
) (models.PasswordRecoveryRequest, error) {
	span := sentry.StartSpan(inputCtx, "Service: PasswordRecoveryRequests.Create")
	defer func() {
		span.Finish()
	}()
	ctx := span.Context()

	passwordRecoveryRequest := models.PasswordRecoveryRequest{
		ID:                uuid.New().String(),
		Email:             params.Email,
		ConfirmationToken: uuid.New().String(),
	}

	existingUser, err := p.repository.Ydb.Users.GetByEmail(ctx, passwordRecoveryRequest.Email)

	if err != nil {
		sentry.CaptureException(err)
		return passwordRecoveryRequest, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestYdbStarterFindUser,
			Description: err.Error(),
		}
	}

	if existingUser.ID == "" {
		return passwordRecoveryRequest, nil
	}

	existingPasswordRecoveryRequest, err :=
		p.repository.Ydb.PasswordRecoveryRequests.GetByEmail(ctx, passwordRecoveryRequest.Email)

	if err != nil {
		sentry.CaptureException(err)
		return passwordRecoveryRequest, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestYdbStarterCreatePasswordRecoveryRequest,
			Description: err.Error(),
		}
	}

	if existingPasswordRecoveryRequest.ID != "" {
		passwordRecoveryRequest = existingPasswordRecoveryRequest
	} else {
		passwordRecoveryRequest, err = p.repository.Ydb.PasswordRecoveryRequests.Create(ctx, passwordRecoveryRequest)
	}

	if err != nil {
		sentry.CaptureException(err)
		return passwordRecoveryRequest, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestYdbStarterCreatePasswordRecoveryRequest,
			Description: err.Error(),
		}
	}

	passwordRecoveryRequestEmail := fmt.Sprintf(
		PasswordRecoveryRequestEmailTemplate,
		p.config.FrontURL,
		p.config.FrontConfirmationPasswordRecoveryRequestURL,
		passwordRecoveryRequest.ConfirmationToken,
	)

	err = p.repository.Ycp.SendEmail(
		ctx,
		passwordRecoveryRequest.Email,
		"Восстановление пароля на "+p.config.FrontURL,
		passwordRecoveryRequestEmail,
	)

	if err != nil {
		sentry.CaptureException(err)
		return passwordRecoveryRequest, &models.Error{
			Code:        models.ErrorCodePasswordRecoveryRequestSMTPSendEmail,
			Description: err.Error(),
		}
	}

	return passwordRecoveryRequest, nil
}
