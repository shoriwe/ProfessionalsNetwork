package Contractor

import (
	"encoding/json"
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"log"
	"net/http"
)

func handleRequestTeamsOwned(contractorID *int) []byte {
	rows, connectionError := SQLDatabase.QuerySelectConnection(Queries.GetOwnedTeamsByProfessional, *contractorID)
	if connectionError == nil {
		results := make([]map[string]interface{}, 0)
		for rows.Next() {
			id := new(int)
			ownerID := new(int)
			teamName := new(string)
			scanError := rows.Scan(id, ownerID, teamName)
			if scanError == nil {
				results = append(results, map[string]interface{}{"id": *id, "ownerID": *ownerID, "name": *teamName})
			}
		}
		marshalResults, marshalError := json.Marshal(results)
		if marshalError == nil {
			return marshalResults
		}
		log.Print(marshalError)
		return []byte("{\"Error\":\"Something goes wrong\"}")
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"Something goes wrong\"}")
}

func RequestTeamsOwned(writer http.ResponseWriter, request *http.Request) {
	requestTeamsOwnedForm := new(RequestsForms.RequestTeamsOwnedForm)
	if contractorID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, requestTeamsOwnedForm, true, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(handleRequestTeamsOwned(&contractorID))
	}
}
