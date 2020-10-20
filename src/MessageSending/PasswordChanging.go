package MessageSending

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"log"
	"net/http"
	"net/url"
)

func SendEmailPasswordHasChanged(accountID *int, username *string) {
	_, email, _ := AccountCore.GetContactInformation(accountID)
	formData := url.Values{
		"APIKey":       {APIKey},
		"Subject":      {"Password changed"},
		"Message":      {"We want to inform you that you or someone else changed the password of your account"},
		"EmailAddress": {email},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/email", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}

func SendEmailResetURL(usernameHash *string, resetKey *string, accountID *int) {
	_, email, _ := AccountCore.GetContactInformation(accountID)
	formData := url.Values{
		"APIKey":       {APIKey},
		"Subject":      {"Password Reset"},
		"Message":      {"<html>You can reset your password with this <a href=\"" + ResetPasswordURL + "?username-hash=" + *usernameHash + "&" + "reset-key=" + *resetKey + "\">URL</a></html>"},
		"EmailAddress": {email},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/email", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}
