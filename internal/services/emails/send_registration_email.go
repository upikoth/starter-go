package emails

import (
	"context"
	"fmt"

	"github.com/upikoth/starter-go/internal/pkg/tracing"
	"go.opentelemetry.io/otel"
)

var registrationEmailTemplate = `
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

func (e *Emails) SendRegistrationEmail(
	inputCtx context.Context,
	email string,
	token string,
) error {
	tracer := otel.Tracer(tracing.GetServiceTraceName())
	ctx, span := tracer.Start(inputCtx, tracing.GetServiceTraceName())
	defer span.End()

	emailMessage := fmt.Sprintf(
		registrationEmailTemplate,
		e.config.FrontURL,
		e.config.FrontConfirmationRegistrationURL,
		token,
	)

	err := e.repositories.ycp.SendEmail(
		ctx,
		email,
		fmt.Sprintf("Регистрация на %s", e.config.FrontURL),
		emailMessage,
	)

	if err != nil {
		tracing.HandleError(span, err)
		return err
	}

	return nil
}
