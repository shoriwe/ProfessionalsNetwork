package Contractor

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ProfessionalCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase/Functionalities"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func acceptProfessionalInTeamBackend(contractorID *int, acceptApplication *bool, applicationKey *string) []byte {
	if Functionalities.IsPendingAppliance(applicationKey) {
		if *acceptApplication {
			rawUserAppliance := Functionalities.GetPendingAppliance(applicationKey)
			userAppliance := new(Templates.UserAppliance)
			scanError := rawUserAppliance.Scan(userAppliance)
			if scanError == nil {
				if *contractorID == userAppliance.ContractorID {
					if !ContractorCore.ProfessionalIsInTeam(&userAppliance.ContractorID,
						&userAppliance.ProfessionalID, &userAppliance.TeamName, &userAppliance.TeamID) {
						if ProfessionalCore.AddProfessionalToTeam(&userAppliance.ProfessionalID, &userAppliance.ContractorID, &userAppliance.TeamID, &userAppliance.TeamName) {
							Functionalities.DeletePendingAppliance(applicationKey)
							return []byte("{\"Success\":\"You are now part of the team\"}")
						}
						return []byte("{\"Error\":\"Something goes wrong, please try again\"}")
					}
					Functionalities.DeletePendingAppliance(applicationKey)
					return []byte("{\"Error\":\"Looks like you where recently added to this team\"}")
				}
				return []byte("{\"Error\":\"You are not owner of this team\"}")
			}
			return []byte("{\"Error\":\"Something Wrong occurred\"}")
		}
		Functionalities.DeletePendingAppliance(applicationKey)
		return []byte("{\"Success\":\"Invitation declined\"}")
	}
	return []byte("{\"Error\":\"Invalid invitation key\"}")
}

func AcceptProfessionalInTeam(writer http.ResponseWriter, request *http.Request) {
	acceptProfessionalApplicationForm := new(RequestsForms.AcceptProfessionalApplicationForm)
	if contractorID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, acceptProfessionalApplicationForm, true, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(acceptProfessionalInTeamBackend(&contractorID, &acceptProfessionalApplicationForm.Accept, &acceptProfessionalApplicationForm.ApplicationKey))
	}
}
