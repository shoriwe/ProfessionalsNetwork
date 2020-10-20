package Neo4jDatabase

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

var Database neo4j.Driver

type Neo4jFilterFunction func(result *neo4j.Result) interface{}

func ExecuteQuery(session neo4j.Session, query string, queryArguments map[string]interface{}, filterFunction Neo4jFilterFunction) interface{} {
	transactionResult, transactionError := session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			queryResult, queryError := transaction.Run(query, queryArguments)
			if queryError != nil {
				return nil, queryError
			}
			return queryResult, queryError
		})
	if transactionError == nil {
		result := transactionResult.(neo4j.Result)
		return filterFunction(&result)
	}
	log.Print(transactionError)
	return nil
}

func CloseSession(session neo4j.Session) {
	if session != nil {
		session.Close()
	}
}

func Connect(databaseUri string, neo4jUsername string, neo4jPassword string) {
	Database, _ = neo4j.NewDriver(databaseUri, neo4j.BasicAuth(neo4jUsername, neo4jPassword, ""), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	session, _ := Database.NewSession(neo4j.SessionConfig{})
	_, connectionError := session.Run("MATCH 1 RETURN 1", nil)
	if connectionError != nil {
		log.Fatal(connectionError)
	}
}
