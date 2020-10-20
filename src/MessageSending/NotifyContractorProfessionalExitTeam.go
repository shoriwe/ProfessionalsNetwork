package MessageSending

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ProfessionalCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"log"
	"net/http"
	"net/url"
)

func SendEmailProfessionalExitFromTeam(contractorID *int, professionalID *int, teamName *string) {
	_, contractorEmail, _ := AccountCore.GetContactInformation(contractorID)
	professionalName := ProfessionalCore.GetProfessionalName(professionalID)
	formData := url.Values{
		"Subject":      {"Professional exit from a team"},
		"APIKey":       {APIKey},
		"Message":      {"<html>" + professionalName + " Has left your " + *teamName + " Team"},
		"EmailAddress": {contractorEmail},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/email", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}
