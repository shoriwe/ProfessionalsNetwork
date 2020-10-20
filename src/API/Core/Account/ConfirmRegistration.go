package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler/TemplateValidation"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func ConfirmRegistration(writer http.ResponseWriter, request *http.Request) {
	emailPhoneNumberValidationForm := new(RequestsForms.EmailPhoneNumberValidationForm)
	if _, _, _, _, isRequestValid := SessionHandling.RequestHandler(writer, request, emailPhoneNumberValidationForm, true); isRequestValid {
		found, message := TemplateValidation.IsEmailPhoneNumberValidationFormValid(emailPhoneNumberValidationForm)
		if found {
			if AccountCore.ConfirmAccountRegisterBackend(&emailPhoneNumberValidationForm.UsernameHash, &emailPhoneNumberValidationForm.EmailKey, &emailPhoneNumberValidationForm.PhoneKey) {
				message = []byte("{\"Success\":\"User successfully registered\"}")
			} else {
				message = []byte("{\"Error\":\"The EmailKey or PhoneKey are incorrect\"}")
			}
		}
		_, _ = writer.Write(message)
	}
}
