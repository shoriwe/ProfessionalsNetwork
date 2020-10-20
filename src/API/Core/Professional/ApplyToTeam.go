package Professional

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

func applyToTeamBackend(professionalID *int, professionalUsername *string, teamName *string, contractorID *int) []byte {
	teamID := AccountCore.TeamAlreadyExists(contractorID, teamName)
	if teamID != -1 {
		if !ContractorCore.ProfessionalIsInTeam(contractorID, professionalID, teamName, &teamID) { // Something to check if the user is already in the team
			hashHandler := sha3.New256()
			hashHandler.Write([]byte(*professionalUsername))
			hashHandler.Write([]byte(*teamName))
			hashHandler.Write([]byte(strconv.Itoa(*contractorID)))
			hashHandler.Write([]byte(strconv.Itoa(teamID)))
			applicationKey := hex.EncodeToString(hashHandler.Sum(nil))
			if !Functionalities.IsPendingAppliance(&applicationKey) { // Something to check if user has an already pending invitation for this team
				userAppliance := new(Templates.UserAppliance)
				userAppliance.TeamName = *teamName
				userAppliance.Username = *professionalUsername
				userAppliance.ProfessionalID = *professionalID
				userAppliance.ContractorID = *contractorID
				userAppliance.TeamID = teamID

				connection := RedisDatabase.GetPendingApplianceConnection()
				defer RedisDatabase.CloseConnection(connection)
				Functionalities.SetPendingApplianceKeyValue(connection, &applicationKey, userAppliance, 24*time.Hour*7)
				go MessageSending.SendEmailAppliance(contractorID, professionalID, teamName, &applicationKey)
				return []byte("{\"Success\":\"You have sent your application form to this team\"}")
			}
			return []byte("{\"Error\":\"The user already has a pending application form with this team\"}")
		}
		return []byte("{\"Error\":\"User is already in this team\"}")
	}
	return []byte("{\"Error\":\"Team doesn't exists\"}")
}

func ApplyToTeam(writer http.ResponseWriter, request *http.Request) {
	applyToTeamForm := new(RequestsForms.ApplyToTeamForm)
	if professionalID, professionalUsername, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, applyToTeamForm, true, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(applyToTeamBackend(&professionalID, &professionalUsername, &applyToTeamForm.TeamName, &applyToTeamForm.ContractorID))
	}
}
