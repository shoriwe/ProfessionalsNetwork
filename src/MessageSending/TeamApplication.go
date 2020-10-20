package MessageSending

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func SendEmailAppliance(contractorID *int, professionalID *int, teamName *string, applicationKey *string) {
	_, contractorEmail, _ := AccountCore.GetContactInformation(contractorID)
	_, professionalEmail, _ := AccountCore.GetContactInformation(professionalID)
	// Link of professional profile
	// Link for accept or decline professional
	// Send professional email
	formData := url.Values{
		"APIKey":       {APIKey},
		"Subject":      {"Team applicant"},
		"Message":      {"<html>Someone want to work in " + *teamName + " You can check his <a href=\"" + ProfessionalProfile + "?add=true&professional-id=" + strconv.Itoa(*professionalID) + "&team-name=" + *teamName + "\">Profile</a><br> To them contact him directly to his/her email " + professionalEmail + "<br>You can accept/reject him by acceding <a href=\"" + InvitationURL + "?application-key=" + *applicationKey + "\">URL</a></html>"},
		"EmailAddress": {contractorEmail},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/email", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}
