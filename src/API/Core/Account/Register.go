package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler/TemplateValidation"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"github.com/shoriwe/ProNet/src/MessageSending"
	"net/http"
)

func registerBackend(registrationForm *RequestsForms.RegistrationForm) []byte {
	isValid, validationError := TemplateValidation.IsRegistrationFormValid(registrationForm)
	if isValid {
		usernameHash, emailKey, phoneKey, userInQueue := AccountCore.PrepareRegistrationData(&registrationForm.Name, &registrationForm.Username, &registrationForm.Password, &registrationForm.Email, &registrationForm.CountryCode, &registrationForm.PhoneNumber, SQLDatabase.AccountTypesByName[registrationForm.AccountType])
		if userInQueue == nil {
			go MessageSending.SendEmailValidationForRegistration(&usernameHash, &registrationForm.Email, &emailKey)
			go MessageSending.SendSMSValidationForRegistration(&registrationForm.CountryCode, &registrationForm.PhoneNumber, &phoneKey)
			return []byte("{\"Success\": \"Waiting for confirmation of email and phone number\"}")
		}
		return []byte("{\"Error\":\"" + userInQueue.Error() + "\"}")
	}
	return []byte(validationError)
}

func Register(writer http.ResponseWriter, request *http.Request) {
	registrationForm := new(RequestsForms.RegistrationForm)
	if _, _, _, _, isRequestValid := SessionHandling.RequestHandler(writer, request, registrationForm, true); isRequestValid {
		_, _ = writer.Write(registerBackend(registrationForm))
	}
}
