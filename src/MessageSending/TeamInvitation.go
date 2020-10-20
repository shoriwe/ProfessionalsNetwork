package MessageSending

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"log"
	"net/http"
	"net/url"
)

func SendEmailInvitation(contractorID *int, teamName *string, professionalID *int, invitationKey *string) {
	_, contractorEmail, _ := AccountCore.GetContactInformation(contractorID)
	_, professionalEmail, _ := AccountCore.GetContactInformation(professionalID)
	contractorName := ContractorCore.GetContractorName(contractorID)
	// Link of professional profile
	// Link for accept or decline professional
	// Send contractor details
	formData := url.Values{
		"APIKey":       {APIKey},
		"Subject":      {"Team Invitation"},
		"Message":      {"<html>" + contractorName + " want to invite you work in his team called " + *teamName + "<br>You can contact him directly to his email for more details " + contractorEmail + "<br>You can accept/decline this invitation with this <a href=\"" + InvitationURL + "?invitation-key=" + *invitationKey + "\">URL</a></html>"},
		"EmailAddress": {professionalEmail},
	}
	_, connectionError := http.PostForm(MessagingAPIURI+"/send/email", formData)
	if connectionError != nil {
		log.Print(connectionError)
	}
}
