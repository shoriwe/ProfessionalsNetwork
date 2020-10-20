package ProfessionalCore

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
	"log"
	"strings"
)

func searchTeamsFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		record := (*result).Record()
		rawTeams, _ := record.Get("teams")
		teams := make([]interface{}, 0)
		for _, team := range rawTeams.([]interface{}) {
			teams = append(teams, team.(neo4j.Node).Props())
		}
		marshalTeams, marshalError := json.Marshal(teams)
		if marshalError == nil {
			return marshalTeams
		}
	}
	return []byte("{\"teams\":[]}")
}

func handleSearchTeams(professionalID *int, searchQuery *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"SearchQuery":    strings.ToLower(*searchQuery),
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.SearchTeamsQuery, data, searchTeamsFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Could not dislike this skill as some error occurred\"}")
}

func SearchTeamsBackend(professionalID *int, searchQuery *string) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleSearchTeams(professionalID, searchQuery, session)
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
