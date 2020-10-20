package Contractor

import (
	"encoding/hex"
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase/Functionalities"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"github.com/shoriwe/ProNet/src/MessageSending"
	"golang.org/x/crypto/sha3"
	"net/http"
	"strconv"
	"time"
)

func inviteProfessionalToTeamBackend(contractorID *int, contractorUsername *string, professionalID *int, teamName *string) []byte {
	if idValid, professionalUsername, accountType := AccountCore.UserIDExists(professionalID); idValid && accountType == SQLDatabase.ProfessionalAccount {
		if teamID := AccountCore.TeamAlreadyExists(contractorID, teamName); teamID != -1 {
			if !ContractorCore.ProfessionalIsInTeam(contractorID, professionalID, teamName, &teamID) { // Something to check if the user is already in the team
				hashHandler := sha3.New256()
				hashHandler.Write([]byte(professionalUsername))
				hashHandler.Write([]byte(*contractorUsername))
				hashHandler.Write([]byte(*teamName))
				hashHandler.Write([]byte(strconv.Itoa(*contractorID)))
				hashHandler.Write([]byte(strconv.Itoa(teamID)))
				invitationKey := hex.EncodeToString(hashHandler.Sum(nil))
				if !Functionalities.IsPendingInvitation(&invitationKey) { // Something to check if user has an already pending invitation for this team
					userInvitation := new(Templates.UserInvitation)
					userInvitation.TeamName = *teamName
					userInvitation.Username = professionalUsername
					userInvitation.ProfessionalID = *professionalID
					userInvitation.ContractorID = *contractorID
					userInvitation.TeamID = teamID

					connection := RedisDatabase.GetPendingInvitationConnection()
					defer RedisDatabase.CloseConnection(connection)
					Functionalities.SetPendingInvitationKeyValue(connection, &invitationKey, userInvitation, 24*time.Hour*7)
					go MessageSending.SendEmailInvitation(contractorID, teamName, professionalID, &invitationKey)
					return []byte("{\"Success\":\"Professional was invited to your team\"}")
				}
				return []byte("{\"Error\":\"The user already have a pending invitation for this team\"}")
			}
			return []byte("{\"Error\":\"User is already in this team\"}")
		}
		return []byte("{\"Error\":\"Team doesn't exists\"}")
	}
	return []byte("{\"Error\":\"Username doesn't exists or it's another contractor\"}")
}

func InviteProfessionalToTeam(writer http.ResponseWriter, request *http.Request) {
	inviteProfessionalToTeam := new(RequestsForms.InviteProfessionalToTeamForm)
	if contractorID, contractorUsername, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, inviteProfessionalToTeam, true, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(inviteProfessionalToTeamBackend(&contractorID, &contractorUsername, &inviteProfessionalToTeam.ProfessionalID, &inviteProfessionalToTeam.TeamName))
	}
}
