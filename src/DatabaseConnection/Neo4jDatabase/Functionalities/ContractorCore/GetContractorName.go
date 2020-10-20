package ContractorCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Contractors"
)

func getContractorNameFilter(result *neo4j.Result) interface{} {
	if (*result).Next() {
		record := (*result).Record()
		name, found := record.Get("name_")
		if found {
			return name.(string)
		}
	}
	return ""
}

func handleGetContractorName(contractorID *int, session neo4j.Session) string {
	data := map[string]interface{}{
		"ContractorID": *contractorID,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Contractors.GetContractorNameQuery, data, getContractorNameFilter)
	if result != nil {
		return result.(string)
	}
	return ""
}

func GetContractorName(contractorID *int) string {
	session, sessionCreationError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if sessionCreationError == nil {
		defer session.Close()
		return handleGetContractorName(contractorID, session)
	}
	return ""
}
