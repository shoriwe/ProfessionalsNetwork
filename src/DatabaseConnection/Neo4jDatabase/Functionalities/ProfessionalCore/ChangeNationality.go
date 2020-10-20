package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
	"log"
)

func changeNationalityFilter(result *neo4j.Result) interface{} {
	if (*result).(neo4j.Result).Next() {
		return []byte("{\"Success\":\"Nationality of professional updated successfully\"}")
	}
	return nil
}

func handleChangeNationality(professionalID *int, nationality *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Nationality":    *nationality,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.ChangeProfessionalNationalityQuery, data, changeNationalityFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Something goes wrong\"}")
}

func ChangeNationalityBackend(professionalID *int, nationality *string) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleChangeNationality(professionalID, nationality, session)
	}
	log.Print(connectionError)
	return []byte("{}")
}
