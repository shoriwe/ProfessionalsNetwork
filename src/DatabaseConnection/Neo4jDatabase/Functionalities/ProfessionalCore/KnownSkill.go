package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
	"log"
)

func professionalKnowsSkillFilter(result *neo4j.Result) interface{} {
	return (*result).Next()
}

func professionalKnowsSkill(professionalID *int, skillName *string, session neo4j.Session) bool {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Skill":          *skillName,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.ProfessionalKnowsSkillQuery, data, professionalKnowsSkillFilter)
	if result != nil {
		return result.(bool)
	}
	return false
}

func knownSkillFilter(_ *neo4j.Result) interface{} {
	return []byte("{\"Success\":\"Successfully known skill\"}")
}

func knownSkill(professionalID *int, skillName *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Skill":          *skillName,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.KnownSkillQuery, data, knownSkillFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Could not know this skill (some error occurred)\"}")
}

func unknownSkillFilter(_ *neo4j.Result) interface{} {
	return []byte("{\"Success\":\"Successfully unknown skill\"}")
}

func unknownSkill(professionalID *int, skillName *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Skill":          *skillName,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.UnknownSkillQuery, data, unknownSkillFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Could not like this skill as some error occurred\"}")
}

func handleUnknownSkill(professionalID *int, skillName *string, session neo4j.Session) []byte {
	if AccountCore.CheckIfSkillExists(skillName, session) {
		if professionalKnowsSkill(professionalID, skillName, session) {
			return unknownSkill(professionalID, skillName, session)
		}
		return []byte("{\"Error\":\"The professional doesn't know  this skill\"}")
	}
	return []byte("{\"Error\":\"Skill doesn't exists\"}")
}

func handleKnownSkill(professionalID *int, skillName *string, session neo4j.Session) []byte {
	if AccountCore.CheckIfSkillExists(skillName, session) {
		if !professionalKnowsSkill(professionalID, skillName, session) {
			return knownSkill(professionalID, skillName, session)
		}
		return []byte("{\"Error\":\"The professional already know this skill\"}")
	}
	return []byte("{\"Error\":\"Skill doesn't exists\"}")
}

func KnownSkillBackend(professionalID *int, skillName *string, know *bool) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		if *know {
			return handleKnownSkill(professionalID, skillName, session)
		}
		return handleUnknownSkill(professionalID, skillName, session)
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
