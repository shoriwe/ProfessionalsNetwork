package AccountCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
)

func createProfessionalNodeFilter(_ *neo4j.Result) interface{} {
	return true
}

func handleCreateProfessionalNode(professionalName *string, professionalID *int, session neo4j.Session) bool {
	data := map[string]interface{}{
		"ProfessionalName": *professionalName,
		"ProfessionalID":   *professionalID,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.CreateProfessionalNodeQuery, data, createProfessionalNodeFilter)
	if result != nil {
		return result.(bool)
	}
	return false
}

func CreateProfessionalNode(professionalName *string, professionalID *int) bool {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleCreateProfessionalNode(professionalName, professionalID, session)
	}
	return false
}
