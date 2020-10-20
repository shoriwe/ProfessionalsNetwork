package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"github.com/shoriwe/ProNet/src/MessageSending"
	"net/http"
)

func changePasswordBackend(userID *int, username *string, oldPassword *string, newPassword *string, newPasswordConfirmation *string) []byte {
	if *newPassword == *newPasswordConfirmation {
		if InputHandler.IsPasswordValid(newPassword) {
			if *oldPassword != *newPassword {
				if AccountCore.UsernameAndPasswordExists(username, oldPassword) {
					updated, message := AccountCore.UpdatePassword(userID, username, newPassword)
					if updated {
						go MessageSending.SendEmailPasswordHasChanged(userID, username)
					}
					return message
				} else {
					return []byte("{\"Error\":\"Looks like the old password is not correct\"}")
				}
			} else {
				return []byte("{\"Error\":\"NewPassword password can't be equal to the old one\"}")
			}
		} else {
			return []byte("{\"Error\":\"The new password is not strong enough\"}")
		}
	}
	return []byte("{\"Error\":\"Looks like the new password confirmation doesnt match\"}")
}

func ChangePassword(writer http.ResponseWriter, request *http.Request) {
	changePasswordForm := new(RequestsForms.ChangeOldNewPasswordForm)
	if userID, username, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, changePasswordForm, true, SQLDatabase.ProfessionalAccount, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(changePasswordBackend(&userID, &username, &changePasswordForm.OldPassword, &changePasswordForm.NewPassword, &changePasswordForm.NewPasswordConfirmation))
	}
}
