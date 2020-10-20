package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
	"log"
)

func likedSkillsOnlyFilter(_ *neo4j.Result) interface{} {
	return []byte("{\"Success\":\"Successfully changed liked only\"}")
}

func changeLikedSkillsOnly(professionalID *int, likedOnly *bool, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"LikedOnly":      *likedOnly,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.ChangeProfessionalLikedOnlyQuery, data, likedSkillsOnlyFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Could not change liked only\"}")
}

func ChangeLikeOnlySkillsBackend(professionalID *int, likedOnly *bool) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return changeLikedSkillsOnly(professionalID, likedOnly, session)
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
