package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
	"log"
)

func changeGenderFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		return []byte("{\"Success\":\"Gender changed successfully\"}")
	}
	return nil
}

func handleChangeGender(professionalID *int, genderName *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Gender":         *genderName,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.ChangeProfessionalGenderQuery, data, changeGenderFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
func ChangeGenderBackend(professionalID *int, genderName *string) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleChangeGender(professionalID, genderName, session)
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
