package ContractorCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Contractors"
)

func professionalIsInTeamFilter(result *neo4j.Result) interface{} {
	return (*result).Next()
}

func handleProfessionalIsInTeam(contractorID *int, professionalID *int, teamName *string, teamID *int, session neo4j.Session) bool {
	data := map[string]interface{}{
		"ContractorID":   *contractorID,
		"ProfessionalID": *professionalID,
		"TeamName":       *teamName,
		"TeamID":         *teamID,
	}
	return Neo4jDatabase.ExecuteQuery(session, Contractors.CheckProfessionalIsInTeamQuery, data, professionalIsInTeamFilter).(bool)
}

func ProfessionalIsInTeam(contractorID *int, professionalID *int, teamName *string, teamID *int) bool {
	session, sessionCreationError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if sessionCreationError == nil {
		defer session.Close()
		return handleProfessionalIsInTeam(contractorID, professionalID, teamName, teamID, session)
	}
	return false
}
