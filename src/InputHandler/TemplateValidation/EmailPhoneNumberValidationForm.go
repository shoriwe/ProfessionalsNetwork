package TemplateValidation

import (
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
)

func IsEmailPhoneNumberValidationFormValid(registrationValidationForm *RequestsForms.EmailPhoneNumberValidationForm) (bool, []byte) {
	if InputHandler.IsUsernameHashValid(&registrationValidationForm.UsernameHash) {
		if InputHandler.IsAValidHex(&registrationValidationForm.EmailKey, 64) {
			if InputHandler.IsAValidHex(&registrationValidationForm.PhoneKey, 12) {
				return true, []byte{}
			} else {
				return false, []byte("{\"Error\":\"The PhoneKey can't be empty\"}")
			}
		} else {
			return false, []byte("{\"Error\":\"The EmailKey can't be empty\"}")
		}
	} else {
		return false, []byte("{\"Error\":\"The Username can't be empty\"}")
	}
}
