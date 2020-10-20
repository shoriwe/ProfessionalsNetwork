package AccountCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/General"
)

func removeProfessionalFromTeamFilter(_ *neo4j.Result) interface{} {
	return nil
}

func handleRemoveProfessionalFromTeam(contractorID *int, professionalID *int, teamName *string, teamID *int, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ContractorID":   *contractorID,
		"ProfessionalID": *professionalID,
		"TeamName":       *teamName,
		"TeamID":         *teamID,
	}
	Neo4jDatabase.ExecuteQuery(session, General.RemoveProfessionalFromTeamQuery, data, removeProfessionalFromTeamFilter)
	return []byte("{\"Success\": \"Professional Removed from team\"}")
}

func RemoveProfessionalFromTeamBackend(contractorID *int, professionalID *int, teamName *string, teamID *int) (bool, []byte) {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		message := handleRemoveProfessionalFromTeam(contractorID, professionalID, teamName, teamID, session)
		return true, message
	}
	return false, []byte("{\"Error\":\"Something goes wrong\"}")
}
