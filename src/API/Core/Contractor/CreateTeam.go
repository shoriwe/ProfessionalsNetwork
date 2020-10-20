package Contractor

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func CreateTeam(writer http.ResponseWriter, request *http.Request) {
	createTeamForm := new(RequestsForms.CreateTeamForm)
	if contractorID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, createTeamForm, true, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(ContractorCore.CreateTeamBackend(&contractorID, &createTeamForm.TeamName))
	}
}
