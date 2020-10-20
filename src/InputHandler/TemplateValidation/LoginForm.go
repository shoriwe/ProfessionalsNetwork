package TemplateValidation

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
)

func IsLoginFormValid(loginForm *RequestsForms.LoginForm, isAdmin bool) (bool, []byte) {
	if InputHandler.IsUsernameValid(&loginForm.Username) {
		if InputHandler.IsPasswordValid(&loginForm.Password) {
			userExists, _, accountType := AccountCore.UsernameExists(&loginForm.Username)
			if userExists {
				if isAdmin == (accountType == SQLDatabase.AdministratorAccount) {
					if AccountCore.UsernameAndPasswordExists(&loginForm.Username, &loginForm.Password) {
						return true, []byte{}
					}
				}
			}
			return false, []byte("{\"Error\":\"Username or password incorrect\"}")
		} else {
			return false, []byte("{\"Error\":\"Password contains invalid chars or doesn't satisfy the minimum rules\"}")
		}
	} else {
		return false, []byte("{\"Error\":\"Username contains invalid chars or is empty\"}")
	}
}
