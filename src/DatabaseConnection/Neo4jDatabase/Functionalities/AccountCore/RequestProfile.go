package AccountCore

import (
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/General"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"log"
	"strconv"
)

func requestRawProfileFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		record := (*result).Record()
		account, found := record.Get("account")
		if found {
			return account.(neo4j.Node).Props()
		}
	}
	return nil
}

func handleRequestRawProfile(accountID *int, session neo4j.Session) interface{} {
	data := map[string]interface{}{
		"AccountID": *accountID,
	}
	return Neo4jDatabase.ExecuteQuery(session, General.GetAccountProfileNodeByIDQuery, data, requestRawProfileFilter)
}

func requestProfileSubNodesFilter(result *neo4j.Result) interface{} {
	subNodes := make([]interface{}, 0)
	if (*result).Next() {
		record := (*result).Record()
		results, _ := record.Get("results")
		for _, resultEntry := range results.([]interface{}) {
			subNodes = append(subNodes, []interface{}{resultEntry.([]interface{})[0].(neo4j.Node).Props(), resultEntry.([]interface{})[1], resultEntry.([]interface{})[2]})
		}
	}
	return subNodes
}

func handleRequestProfileSubNodes(accountID *int, session neo4j.Session) []interface{} {
	data := map[string]interface{}{
		"AccountID": *accountID,
	}
	return Neo4jDatabase.ExecuteQuery(session, General.GetAccountSubNodesByIDQuery, data, requestProfileSubNodesFilter).([]interface{})
}

func handleRequestProfile(accountID *int, accountType *int, session neo4j.Session) []byte {
	temporaryProfile := handleRequestRawProfile(accountID, session)
	if temporaryProfile != nil {
		profile := temporaryProfile.(map[string]interface{})
		profile["type"] = SQLDatabase.AccountTypesByNumber[*accountType]
		subNodes := handleRequestProfileSubNodes(accountID, session)
		links := make([]interface{}, 0)
		nodes := make([]interface{}, 0)
		nodes = append(nodes, profile)
		for index := range subNodes {
			subNodeEntry := subNodes[index].([]interface{})
			subNode := subNodeEntry[0].(map[string]interface{})
			subNode["type"] = subNodeEntry[1]
			subNode["labels"] = subNodeEntry[2]
			nodeTypeCharID := subNode["type"].(string)[:1]
			subNode["id"] = nodeTypeCharID + strconv.Itoa(int(subNode["id"].(int64)))
			nodes = append(nodes, subNode)
			links = append(links, map[string]interface{}{"source": *accountID, "target": subNode["id"]})
		}
		marshalNodesAndLinks, marshalError := json.Marshal(map[string]interface{}{"links": links, "nodes": nodes})
		if marshalError == nil {
			return marshalNodesAndLinks
		}
		log.Print(marshalError)
	}
	return []byte("{\"Error\":\"Something goes wrong\"}")
}

func RequestProfileBackend(accountID *int, accountType int) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleRequestProfile(accountID, &accountType, session)
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"" + connectionError.Error() + "\"}")
}
