package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	AccountCore2 "github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func requestTeamMembersBackend(accountID *int, accountType *int, contractorID *int, teamName *string) []byte {
	teamID := AccountCore.TeamAlreadyExists(contractorID, teamName)
	if teamID != -1 {
		isValidForTheRequest := false
		if *accountType == SQLDatabase.ContractorAccount {
			if *accountID == *contractorID {
				isValidForTheRequest = true
			}
		} else if *accountType == SQLDatabase.ProfessionalAccount {
			isValidForTheRequest = ContractorCore.ProfessionalIsInTeam(contractorID, accountID, teamName, &teamID)
		}
		if isValidForTheRequest {
			return AccountCore2.RequestTeamMembersBackend(contractorID, &teamID, teamName)
		}
		return []byte("{\"Error\":\"You are not part of this team\"}")
	}
	return []byte("{\"Error\":\"The team doesn't exists\"}")
}

func RequestTeamMembers(writer http.ResponseWriter, request *http.Request) {
	requestTeamMembersForm := new(RequestsForms.RequestTeamMembersForm)
	if accountID, _, accountType, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, requestTeamMembersForm, true, SQLDatabase.ContractorAccount, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(requestTeamMembersBackend(&accountID, &accountType, &requestTeamMembersForm.ContractorID, &requestTeamMembersForm.TeamName))
	}
}
