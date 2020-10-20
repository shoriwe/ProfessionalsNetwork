package AccountCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/General"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"log"
)

func changeDescriptionFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		return []byte("{\"Success\":\"Description updated successfully\"}")
	}
	return nil
}

func handleChangeDescription(accountID *int, description *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"AccountID":   *accountID,
		"Description": *description,
	}
	result := Neo4jDatabase.ExecuteQuery(session, General.ChangeAccountDescriptionQuery, data, changeDescriptionFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Could not change the description\"}")
}

func ChangeDescriptionBackend(accountID *int, description *string) []byte {
	if InputHandler.IsDescriptionValid(description) {
		session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
		defer Neo4jDatabase.CloseSession(session)
		if connectionError == nil {
			return handleChangeDescription(accountID, description, session)
		}
		log.Print(connectionError)
		return []byte("{\"Error\":\"Something goes wrong\"}")
	}
	return []byte("{\"Error\":\"Description is invalid or length is longer than 350\"}")
}
