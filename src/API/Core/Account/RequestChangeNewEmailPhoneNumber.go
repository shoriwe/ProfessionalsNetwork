package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"github.com/shoriwe/ProNet/src/MessageSending"
	"log"
	"net/http"
)

func requestChangeEmailPhoneNumberBackend(accountID *int, username *string, emailAddress *string, countryCode *string, phoneNumber *string) []byte {
	if InputHandler.IsEmailValid(emailAddress) {
		if InputHandler.IsCountryCodeValid(countryCode) {
			if InputHandler.IsPhoneNumberValid(phoneNumber) {
				usernameHash, emailKey, phoneKey, preparingError := AccountCore.PrepareChangeEmailPhoneNumberData(accountID, username, emailAddress, countryCode, phoneNumber)
				if preparingError == nil {
					go MessageSending.SendEmailValidationForEmailPhoneNumberChange(&usernameHash, emailAddress, &emailKey)
					go MessageSending.SendSMSValidationForEmailPhoneNumberChange(countryCode, phoneNumber, &phoneKey)
					return []byte("{\"Success\": \"Waiting for confirmation of email and phone number\"}")
				} else {
					log.Print(preparingError)
					return []byte("{\"Error\":\"This user  already has a pending change email and phone number\"}")
				}
			} else {
				return []byte("{\"Error\":\"Phone number has invalid chars\"}")
			}
		} else {
			return []byte("{\"Error\":\"Country code is not valid\"}")
		}
	} else {
		return []byte("{\"Error\":\"The Email has invalid chars\"}")
	}
}

func RequestChangeEmailPhoneNumber(writer http.ResponseWriter, request *http.Request) {
	requestChangeEmailPhoneNumberForm := new(RequestsForms.RequestChangeEmailPhoneNumberForm)
	if accountID, username, _, _, isRequestValid := SessionHandling.RequestHandler(writer, request, requestChangeEmailPhoneNumberForm, true, SQLDatabase.ContractorAccount, SQLDatabase.ProfessionalAccount); isRequestValid {
		_, _ = writer.Write(requestChangeEmailPhoneNumberBackend(&accountID, &username,
			&requestChangeEmailPhoneNumberForm.Email, &requestChangeEmailPhoneNumberForm.CountryCode, &requestChangeEmailPhoneNumberForm.PhoneNumber))
	}
}
