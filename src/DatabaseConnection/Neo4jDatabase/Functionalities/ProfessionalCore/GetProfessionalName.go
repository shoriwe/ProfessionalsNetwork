package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
)

func getProfessionalNameFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		record := (*result).Record()
		name, found := record.Get("name_")
		if found {
			return name
		}
	}
	return ""
}

func handleGetProfessionalName(professionalID *int, session neo4j.Session) string {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.GetProfessionalNameQuery, data, getProfessionalNameFilter)
	if result != nil {
		return result.(string)
	}
	return ""
}

func GetProfessionalName(professionalID *int) string {
	session, sessionCreationError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if sessionCreationError == nil {
		defer session.Close()
		return handleGetProfessionalName(professionalID, session)
	}
	return ""
}
