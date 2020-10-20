package AccountCore

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/General"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"log"
	"strings"
)

func searchLanguagesFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		record := (*result).Record()
		languages, _ := record.Get("languages")
		encodedResult, _ := json.Marshal(map[string]interface{}{"Languages": languages})
		return encodedResult
	}
	return nil
}

func handleSearchLanguages(session neo4j.Session, ownerID *int, objectID *int, searchQuery *string) []byte {
	data := map[string]interface{}{
		"OwnerID":     *ownerID,
		"ObjectID":    *objectID,
		"SearchQuery": strings.ToLower(*searchQuery),
	}
	result := Neo4jDatabase.ExecuteQuery(session, General.SearchLanguagesQuery, data, searchLanguagesFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("[]")
}

func SearchLanguagesBackend(ownerID *int, objectID *int, searchQuery *string) []byte {
	if InputHandler.IsStringValid(searchQuery) {
		session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
		defer Neo4jDatabase.CloseSession(session)
		message := []byte("{\"Error\": \"Connection error\"}")
		if connectionError == nil {
			if IsOwnerOfObject(session, ownerID, objectID) {
				message = handleSearchLanguages(session, ownerID, objectID, searchQuery)
			} else {
				message = []byte("{\"Error\": \"Sorry but the object doesn't exists\"}")
			}
		} else {
			log.Print(connectionError)
		}
		return message
	} else {
		return []byte("{\"Error\": \"Sorry but the submitted string has illegal chars or is empty\"}")
	}
}
