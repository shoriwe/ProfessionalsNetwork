package AccountCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/General"
	"github.com/shoriwe/ProNet/src/InputHandler"
)

func changeProfessionalFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		return []byte("{\"Success\":\"Location updated\"}")
	}
	return nil
}

func handleChangeLocation(accountID *int, locationName *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"AccountID": *accountID,
		"Location":  *locationName,
	}
	result := Neo4jDatabase.ExecuteQuery(session, General.ChangeAccountLocationQuery, data, changeProfessionalFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Could not change the location\"}")
}

func ChangeLocationBackend(accountID *int, locationName *string) []byte {
	var message []byte
	if InputHandler.IsLocationValid(locationName) {
		session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
		defer Neo4jDatabase.CloseSession(session)
		if connectionError == nil {
			message = handleChangeLocation(accountID, locationName, session)
		} else {
			message = []byte("{\"Error\":\"" + connectionError.Error() + "\"}")
		}
	} else {
		message = []byte("{\"Error\":\"The location you provided is not valid\"}")
	}
	return message
}
