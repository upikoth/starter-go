package registrations

import (
	"context"
	"fmt"

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
	context context.Context,
	params models.RegistrationCreateParams,
) (models.Registration, error) {
	registration := models.Registration{
		ID:                uuid.New().String(),
		Email:             params.Email,
		ConfirmationToken: uuid.New().String(),
	}

	registrationEmail := fmt.Sprintf(
		RegistrationEmailTemplate,
		r.config.FrontURL,
		r.config.FrontConfirmationRegistrationURL,
		registration.ConfirmationToken,
	)

	err := r.repository.YcpStarter.SendEmail(
		registration.Email,
		"Регистрация на "+r.config.FrontURL,
		registrationEmail,
	)

	if err != nil {
		return registration, &models.Error{
			Code:        models.ErrorCodeRegistrationSMTPSendEmail,
			Description: err.Error(),
		}
	}

	resRegistration, err := r.repository.YdbStarter.Registrations.Create(context, registration)

	if err != nil {
		return registration, &models.Error{
			Code:        models.ErrorCodeRegistrationYdbStarterCreateEmail,
			Description: err.Error(),
		}
	}

	return resRegistration, nil
}
