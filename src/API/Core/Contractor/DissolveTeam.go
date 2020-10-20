package Contractor

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	ContractorCore2 "github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func handleDissolveTeam(contractorID *int, teamName *string) []byte {
	teamID := AccountCore.TeamAlreadyExists(contractorID, teamName)
	if teamID != -1 {
		if ContractorCore.DissolveTeamBackend(&teamID, contractorID, teamName) {
			if ContractorCore2.DissolveTeam(&teamID, contractorID, teamName) {
				return []byte("{\"Success\":\"Team dissolved successfully\"}")
			}
		}
		return []byte("{\"Error\":\"Something goes wrong, please try again\"}")
	}
	return []byte("{\"Error\":\"Team doesn't exists\"}")
}

func DissolveTeam(writer http.ResponseWriter, request *http.Request) {
	dissolveTeamForm := new(RequestsForms.DissolveTeamForm)
	if contractorID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, dissolveTeamForm, true, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(handleDissolveTeam(&contractorID, &dissolveTeamForm.TeamName))
	}
}
