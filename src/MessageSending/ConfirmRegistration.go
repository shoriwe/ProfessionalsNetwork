package MessageSending

import (
	"log"
	"net/http"
	"net/url"
)

func SendEmailValidationForRegistration(usernameHash *string, email *string, emailKey *string) {
	formData := url.Values{
		"Subject":      {"Account Validation"},
		"Message":      {"<html>Use this <a href=\"" + AccountConfirmationURL + "?username-hash=" + *usernameHash + "&" + "email-key=" + *emailKey + "\">URL</a> to confirm your email and phone number</html>"},
		"EmailAddress": {*email},
		"APIKey":       {APIKey},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/email", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}

func SendSMSValidationForRegistration(countryCode *string, phoneNumber *string, phoneKey *string) {
	formData := url.Values{
		"APIKey":      {APIKey},
		"CountryCode": {*countryCode},
		"PhoneNumber": {*phoneNumber},
		"Message":     {"Your registration code is: " + *phoneKey},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/sms", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}
