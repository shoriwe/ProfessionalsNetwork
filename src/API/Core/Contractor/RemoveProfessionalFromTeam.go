package Contractor

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	AccountCore2 "github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"github.com/shoriwe/ProNet/src/MessageSending"
	"net/http"
)

func RemoveProfessionalFromTeam(writer http.ResponseWriter, request *http.Request) {
	removeProfessionalFromTeamForm := new(RequestsForms.RemoveProfessionalFromTeamForm)
	if contractorID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, removeProfessionalFromTeamForm, true, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		teamID := AccountCore2.TeamAlreadyExists(&contractorID, &removeProfessionalFromTeamForm.TeamName)
		if teamID != -1 {
			if ContractorCore.ProfessionalIsInTeam(&contractorID, &removeProfessionalFromTeamForm.ProfessionalID, &removeProfessionalFromTeamForm.TeamName, &teamID) {
				professionalFired, message := AccountCore.RemoveProfessionalFromTeamBackend(&contractorID, &removeProfessionalFromTeamForm.ProfessionalID, &removeProfessionalFromTeamForm.TeamName, &teamID)
				if professionalFired {
					go MessageSending.SendEmailProfessionalFiredFromTeam(&removeProfessionalFromTeamForm.ProfessionalID, &removeProfessionalFromTeamForm.TeamName)
				}
				_, _ = writer.Write(message)
			} else {
				_, _ = writer.Write([]byte("{\"Error\":\"Professional is not part of this team\"}"))
			}
		} else {
			_, _ = writer.Write([]byte("{\"Error\":\"The team doesn't exists\"}"))
		}
	}
}
