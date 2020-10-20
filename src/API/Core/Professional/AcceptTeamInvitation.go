package Professional

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ProfessionalCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase/Functionalities"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"log"
	"net/http"
)

func acceptTeamInvitationBackend(professionalID *int, invitationKey *string, acceptInvitation *bool) []byte {
	if Functionalities.IsPendingInvitation(invitationKey) {
		if *acceptInvitation {
			rawUserInvitation := Functionalities.GetPendingInvitation(invitationKey)
			userInvitation := new(Templates.UserInvitation)
			scanError := rawUserInvitation.Scan(userInvitation)
			if scanError == nil {
				if *professionalID == userInvitation.ProfessionalID {
					if !ContractorCore.ProfessionalIsInTeam(&userInvitation.ContractorID,
						&userInvitation.ProfessionalID, &userInvitation.TeamName, &userInvitation.TeamID) {
						if ProfessionalCore.AddProfessionalToTeam(&userInvitation.ProfessionalID, &userInvitation.ContractorID, &userInvitation.TeamID, &userInvitation.TeamName) {
							Functionalities.DeletePendingInvitation(invitationKey)
							return []byte("{\"Success\":\"You are now part of the team\"}")
						}
						return []byte("{\"Error\":\"Something goes wrong, please try again\"}")
					}
					Functionalities.DeletePendingInvitation(invitationKey)
					return []byte("{\"Error\":\"Looks like you where recently added to this team\"}")
				}
				return []byte("{\"Error\":\"You was not invited to this team\"}")
			}
			log.Print(scanError)
			return []byte("{\"Error\":\"Something Wrong occurred\"}")
		}
		Functionalities.DeletePendingInvitation(invitationKey)
		return []byte("{\"Success\":\"Invitation declined\"}")

	}
	return []byte("{\"Error\":\"Invalid invitation key\"}")
}

func AcceptTeamInvitation(writer http.ResponseWriter, request *http.Request) {
	acceptTeamInvitationForm := new(RequestsForms.AcceptTeamInvitationForm)
	if professionalID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, acceptTeamInvitationForm, true, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(acceptTeamInvitationBackend(&professionalID, &acceptTeamInvitationForm.InvitationKey, &acceptTeamInvitationForm.Accept))
	}
}
