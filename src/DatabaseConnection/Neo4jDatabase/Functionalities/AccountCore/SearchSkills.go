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

func searchSkillsFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		record := (*result).Record()
		skills, _ := record.Get("skills")
		encodedResult, _ := json.Marshal(map[string]interface{}{"Skills": skills})
		return encodedResult
	}
	return nil
}

func SearchSkills(session neo4j.Session, ownerID *int, objectID *int, searchQuery *string) []byte {
	data := map[string]interface{}{
		"OwnerID":     ownerID,
		"ObjectID":    *objectID,
		"SearchQuery": strings.ToLower(*searchQuery),
	}
	result := Neo4jDatabase.ExecuteQuery(session, General.SearchSkillsQuery, data, searchSkillsFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("[]")
}

func SearchSkillsBackend(ownerID *int, objectID *int, searchQuery *string) []byte {
	if InputHandler.IsStringValid(searchQuery) {
		session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
		defer Neo4jDatabase.CloseSession(session)
		message := []byte("{\"Error\": \"Connection Error\"}")
		if connectionError == nil {
			if IsOwnerOfObject(session, ownerID, objectID) {
				message = SearchSkills(session, ownerID, objectID, searchQuery)
			} else {
				message = []byte("{\"Error\": \"Sorry but the object doesn't exists or you are not the it\"}")
			}
		} else {
			log.Print(connectionError)
		}
		return message
	} else {
		return []byte("{\"Error\": \"Sorry but the submitted string has illegal chars or is empty\"}")
	}
}
