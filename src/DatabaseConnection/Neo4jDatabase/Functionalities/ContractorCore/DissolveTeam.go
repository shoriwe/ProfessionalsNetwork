package ContractorCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Contractors"
)

func dissolveTeamFilter(_ *neo4j.Result) interface{} {
	return nil
}

func handleDissolveTeam(teamID *int, contractorID *int, teamName *string, session neo4j.Session) bool {
	data := map[string]interface{}{
		"TeamID":   *teamID,
		"OwnerID":  *contractorID,
		"TeamName": *teamName,
	}
	Neo4jDatabase.ExecuteQuery(session, Contractors.DissolveTeamQuery, data, dissolveTeamFilter)
	return true
}

func DissolveTeamBackend(teamID *int, contractorID *int, teamName *string) bool {
	session, sessionCreationError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if sessionCreationError == nil {
		defer session.Close()
		return handleDissolveTeam(teamID, contractorID, teamName, session)
	}
	return false
}
