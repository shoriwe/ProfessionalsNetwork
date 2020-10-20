package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
)

func changeRemoteFilter(result *neo4j.Result) interface{} {
	if (*result).(neo4j.Result).Next() {
		return []byte("{\"Success\":\"Remote state changed successfully\"}")
	}
	return nil
}

func handleChangeRemote(professionalID *int, remote *bool, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Remote":         *remote,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.ChangeRemoteQuery, data, changeRemoteFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
func ChangeRemoteBackend(professionalID *int, remote *bool) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleChangeRemote(professionalID, remote, session)
	}
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
