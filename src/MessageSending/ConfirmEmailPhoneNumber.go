package MessageSending

import (
	"log"
	"net/http"
	"net/url"
)

func SendEmailValidationForEmailPhoneNumberChange(usernameHash *string, email *string, emailKey *string) {
	formData := url.Values{
		"Subject":      {"Contact Information Updated Validation"},
		"Message":      {"<html>Use this <a href=\"" + EmailPhoneNumberChangeValidationURL + "?username-hash=" + *usernameHash + "&" + "email-key=" + *emailKey + "\">URL</a> to confirm your email and phone number</html>"},
		"EmailAddress": {*email},
		"APIKey":       {APIKey},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/email", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}

func SendSMSValidationForEmailPhoneNumberChange(countryCode *string, phoneNumber *string, phoneKey *string) {
	formData := url.Values{
		"APIKey":      {APIKey},
		"CountryCode": {*countryCode},
		"PhoneNumber": {*phoneNumber},
		"Message":     {"Your confirmation code is: " + *phoneKey},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/sms", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}
