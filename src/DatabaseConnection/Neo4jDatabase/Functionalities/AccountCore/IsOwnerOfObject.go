package AccountCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/General"
)

func isOwnerOfObjectFilter(result *neo4j.Result) interface{} {
	return (*result).Next()
}

func IsOwnerOfObject(session neo4j.Session, ownerID *int, objectID *int) bool {
	data := map[string]interface{}{
		"OwnerID":  *ownerID,
		"ObjectID": *objectID,
	}
	result := Neo4jDatabase.ExecuteQuery(session, General.CheckIfIsOwnerOfObjectByIDsQuery, data, isOwnerOfObjectFilter)
	return result.(bool)
}
