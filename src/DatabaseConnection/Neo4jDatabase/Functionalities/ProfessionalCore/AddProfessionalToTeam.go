package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
)

func addProfessionalToTeamFilter(_ *neo4j.Result) interface{} {
	return true
}

func handleAddProfessionalToTeam(professionalID *int, contractorID *int, teamID *int, teamName *string, session neo4j.Session) bool {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"ContractorID":   *contractorID,
		"TeamID":         *teamID,
		"TeamName":       *teamName,
	}
	Neo4jDatabase.ExecuteQuery(session, Professionals.AddProfessionalToTeamQuery, data, addProfessionalToTeamFilter)
	return true
}

func AddProfessionalToTeam(professionalID *int, contractorID *int, teamID *int, teamName *string) bool {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleAddProfessionalToTeam(professionalID, contractorID, teamID, teamName, session)
	}
	return false
}
