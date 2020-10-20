package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
	"log"
)

func changeAvailableFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		return []byte("{\"Success\":\"Available state changed successfully\"}")
	}
	return nil
}

func handleChangeAvailable(professionalID *int, available *bool, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Available":      *available,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.ChangeAvailableQuery, data, changeAvailableFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Something goes wrong\"}")
}

func ChangeAvailableBackend(professionalID *int, available *bool) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleChangeAvailable(professionalID, available, session)
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
