package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase/Functionalities"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"github.com/shoriwe/ProNet/src/MessageSending"
	"net/http"
)

func passwordResetBackend(usernameHash *string, resetKey *string, newPassword *string, newPasswordConfirmation *string, ) []byte {
	if *newPassword == *newPasswordConfirmation {
		if InputHandler.IsPasswordValid(newPassword) {
			if InputHandler.IsPasswordValid(newPasswordConfirmation) {
				updated, message, userID, username := AccountCore.ResetPassword(usernameHash, resetKey, newPassword)
				if updated {
					go MessageSending.SendEmailPasswordHasChanged(&userID, &username)
				}
				return message
			} else {
				return []byte("{\"Error\": \"The confirmation password doesn't accomplish the minimum rules\"}")
			}
		} else {
			return []byte("{\"Error\": \"The password doesn't accomplish the minimum rules\"}")
		}
	}
	return []byte("{\"Error\": \"The passwords doesn't match\"}")
}

func PasswordReset(writer http.ResponseWriter, request *http.Request) {
	resetPasswordForm := new(RequestsForms.ResetPasswordForm)
	if _, _, _, _, isRequestValid := SessionHandling.RequestHandler(writer, request, resetPasswordForm, true); isRequestValid {
		if Functionalities.UserResetExists(&resetPasswordForm.UsernameHash) {
			_, _ = writer.Write(passwordResetBackend(&resetPasswordForm.UsernameHash, &resetPasswordForm.ResetKey, &resetPasswordForm.NewPassword, &resetPasswordForm.NewPasswordConfirmation))
		} else {
			_, _ = writer.Write([]byte("{\"Error\": \"Invalid password change key\"}"))
		}
	}
}
