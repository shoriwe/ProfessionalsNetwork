package TemplateValidation

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
)

func IsRegistrationFormValid(registrationForm *RequestsForms.RegistrationForm) (bool, string) {
	if InputHandler.IsNameValid(&registrationForm.Name) {
		if InputHandler.IsUsernameValid(&registrationForm.Username) {
			if InputHandler.IsPasswordValid(&registrationForm.Password) {
				if InputHandler.IsPasswordValid(&registrationForm.PasswordConfirmation) {
					if InputHandler.IsEmailValid(&registrationForm.Email) {
						if InputHandler.IsEmailValid(&registrationForm.EmailConfirmation) {
							if InputHandler.IsCountryCodeValid(&registrationForm.CountryCode) {
								if InputHandler.IsPhoneNumberValid(&registrationForm.PhoneNumber) {
									if InputHandler.IsStringValid(&registrationForm.AccountType) {
										if registrationForm.Password == registrationForm.PasswordConfirmation {
											if registrationForm.Email == registrationForm.EmailConfirmation {
												if SQLDatabase.AccountTypesByName[registrationForm.AccountType] == SQLDatabase.ProfessionalAccount || SQLDatabase.AccountTypesByName[registrationForm.AccountType] == SQLDatabase.ContractorAccount {
													if found, _, _ := AccountCore.UsernameExists(&registrationForm.Username); !found {
														return true, "{\"Success\": \"Your registration request were added to our queue\"}"
													} else {
														return false, "{\"Error\":\"Username already exists\"}"
													}
												} else {
													return false, "{\"Error\":\"Know account type\"}"
												}
											} else {
												return false, "{\"Error\":\"The confirmation email is not equal to the email provided\"}"
											}
										} else {
											return false, "{\"Error\":\"The confirmation password is not equal to the password provided\"}"
										}
									} else {
										return false, "{\"Error\":\"The AccountType field is not valid\"}"
									}
								} else {
									return false, "{\"Error\":\"The Phone number field is not valid\"}"
								}
							} else {
								return false, "{\"Error\":\"The country code is not valid\"}"
							}
						} else {
							return false, "{\"Error\":\"The confirmation email field is not valid\"}"
						}
					} else {
						return false, "{\"Error\":\"The email field is not valid\"}"
					}
				} else {
					return false, "{\"Error\":\"Confirmation password contains invalid chars or doesn't satisfy the minimum rules\"}"
				}
			} else {
				return false, "{\"Error\":\"Password contains invalid chars or doesn't satisfy the minimum rules\"}"
			}
		} else {
			return false, "{\"Error\":\"Username contains invalid chars or is empty\"}"
		}
	} else {
		return false, "{\"Error\":\"Name contains invalid chars or is empty\"}"
	}
}
