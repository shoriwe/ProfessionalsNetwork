package MessageSending

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"log"
	"net/http"
	"net/url"
)

func SendEmailProfessionalFiredFromTeam(professionalID *int, teamName *string) {
	_, professionalEmail, _ := AccountCore.GetContactInformation(professionalID)
	formData := url.Values{
		"APIKey":       {APIKey},
		"Subject":      {"Fired from a team"},
		"Message":      {"Sorry but you was fired from " + *teamName + ", we recommended to check you Dashboard for more information"},
		"EmailAddress": {professionalEmail},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/email", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}
