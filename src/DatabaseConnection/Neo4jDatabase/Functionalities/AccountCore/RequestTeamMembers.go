package AccountCore

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/General"
	"log"
	"strconv"
)

func requestTeamMembersFilter(result *neo4j.Result) interface{} {
	results := make([]interface{}, 0)
	if (*result).Next() {
		record := (*result).Record()
		rawResults, _ := record.Get("results")
		for _, resultEntry := range rawResults.([]interface{}) {
			results = append(results, []interface{}{resultEntry.([]interface{})[0].(neo4j.Node).Props(), resultEntry.([]interface{})[1], resultEntry.([]interface{})[2]})
		}
	}
	return results
}

func handleRequestTeamMembers(contractorID *int, teamID *int, teamName *string, session neo4j.Session) []interface{} {
	data := map[string]interface{}{
		"ContractorID": *contractorID,
		"TeamName":     *teamName,
		"TeamID":       *teamID,
	}
	result := Neo4jDatabase.ExecuteQuery(session, General.GetTeamMembersQuery, data, requestTeamMembersFilter)
	if result != nil {
		return result.([]interface{})
	}
	return []interface{}{}
}

func RequestTeamMembersBackend(contractorID *int, teamID *int, teamName *string) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	if connectionError == nil {
		realTeamID := "T" + strconv.Itoa(*teamID)
		members := handleRequestTeamMembers(contractorID, teamID, teamName, session)
		nodes := []interface{}{map[string]interface{}{"id": realTeamID, "ownerID": *contractorID, "name": *teamName}}
		nodes = append(nodes, members...)
		links := []map[string]interface{}{}
		for _, member := range members {
			links = append(links, map[string]interface{}{"source": member.([]interface{})[0].(map[string]interface{})["id"], "target": realTeamID})
		}
		marshalNodes, marshalError := json.Marshal(map[string]interface{}{"links": links, "nodes": nodes})
		if marshalError == nil {
			return marshalNodes
		}
		return []byte("{\"Error\":\"Something goes wrong, please try again\"}")
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
