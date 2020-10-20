package ContractorCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Contractors"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/InputHandler"
)

func teamAssociationFilter(_ *neo4j.Result) interface{} {
	return []byte("{\"Success\":\"Team Created successfully\"}")
}

func createTeamFilter(_ *neo4j.Result) interface{} {
	return nil
}

func handleTeamAssociation(contractorID *int, teamID *int, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"TeamID":       *teamID,
		"ContractorID": *contractorID,
	}
	return Neo4jDatabase.ExecuteQuery(session, Contractors.ContractorOwnerOfTeamQuery, data, teamAssociationFilter).([]byte)
}

func handleCreateTeam(contractorID *int, teamName *string, session neo4j.Session) []byte {
	if AccountCore.TeamAlreadyExists(contractorID, teamName) == -1 {
		if ContractorCore.CreateTeamEntry(contractorID, teamName) {
			teamID := AccountCore.TeamAlreadyExists(contractorID, teamName)
			data := map[string]interface{}{
				"TeamID":       teamID,
				"ContractorID": *contractorID,
				"TeamName":     *teamName,
			}
			Neo4jDatabase.ExecuteQuery(session, Contractors.CreateTeamNodeQuery, data, createTeamFilter)
			return handleTeamAssociation(contractorID, &teamID, session)

		}
		return []byte("{\"Error\":\"Something goes wrong\"}")
	}
	return []byte("{\"Error\":\"You already have a team  with that name\"}")
}

func CreateTeamBackend(contractorID *int, teamName *string) []byte {
	if InputHandler.IsTeamNameValid(teamName) {
		session, sessionCreationError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
		defer Neo4jDatabase.CloseSession(session)
		if sessionCreationError == nil {
			defer session.Close()
			return handleCreateTeam(contractorID, teamName, session)
		}
		return []byte("{\"Error\":\"Something goes wrong\"}")
	}
	return []byte("{\"Error\":\"The team name has invalid chars or is empty\"}")
}
