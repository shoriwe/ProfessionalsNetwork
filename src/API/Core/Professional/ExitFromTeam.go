package Professional

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

func ExitFromTeam(writer http.ResponseWriter, request *http.Request) {
	exitFromTeamForm := new(RequestsForms.ExitFromTeamForm)
	if professionalID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, exitFromTeamForm, true, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		teamID := AccountCore2.TeamAlreadyExists(&exitFromTeamForm.ContractorID, &exitFromTeamForm.TeamName)
		if teamID != -1 {
			if ContractorCore.ProfessionalIsInTeam(&exitFromTeamForm.ContractorID, &professionalID, &exitFromTeamForm.TeamName, &teamID) {
				exitFromTeam, message := AccountCore.RemoveProfessionalFromTeamBackend(&exitFromTeamForm.ContractorID, &professionalID, &exitFromTeamForm.TeamName, &teamID)
				if exitFromTeam {
					go MessageSending.SendEmailProfessionalExitFromTeam(&exitFromTeamForm.ContractorID, &professionalID, &exitFromTeamForm.TeamName)
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
