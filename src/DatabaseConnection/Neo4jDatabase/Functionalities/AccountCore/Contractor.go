package AccountCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Contractors"
)

func createContractorNodeFilter(_ *neo4j.Result) interface{} {
	return true
}

func handleCreateContractorNode(contractorName *string, contractorID *int, session neo4j.Session) bool {
	data := map[string]interface{}{
		"ContractorName": *contractorName,
		"ContractorID":   *contractorID,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Contractors.CreateContractorNodeQuery, data, createContractorNodeFilter)
	if result != nil {
		return result.(bool)
	}
	return false
}

func CreateContractorNode(contractorName *string, contractorID *int) bool {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleCreateContractorNode(contractorName, contractorID, session)
	}
	return false
}
