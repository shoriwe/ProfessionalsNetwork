package ProfessionalCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/Professionals"
	"log"
)

func professionalLikedSkillFilter(result *neo4j.Result) interface{} {
	return (*result).Next()
}

func professionalLikesSkill(professionalID *int, skillName *string, session neo4j.Session) bool {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Skill":          *skillName,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.ProfessionalLikesSkillQuery, data, professionalLikedSkillFilter)
	if result != nil {
		return result.(bool)
	}
	return false
}

func likeSkillFilter(_ *neo4j.Result) interface{} {
	return []byte("{\"Success\":\"Successfully liked skill\"}")
}

func likeSkill(professionalID *int, skillName *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Skill":          *skillName,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.LikeSkillQuery, data, likeSkillFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Could not like this skill as some error occurred\"}")
}

func dislikeSkillFilter(_ *neo4j.Result) interface{} {
	return []byte("{\"Success\":\"Successfully disliked the skill\"}")
}

func dislikeSkill(professionalID *int, skillName *string, session neo4j.Session) []byte {
	data := map[string]interface{}{
		"ProfessionalID": *professionalID,
		"Skill":          *skillName,
	}
	result := Neo4jDatabase.ExecuteQuery(session, Professionals.DislikeSkillQuery, data, dislikeSkillFilter)
	if result != nil {
		return result.([]byte)
	}
	return []byte("{\"Error\":\"Could not dislike this skill as some error occurred\"}")
}

func handleDislikeSkill(professionalID *int, skillName *string, session neo4j.Session) []byte {
	if AccountCore.CheckIfSkillExists(skillName, session) {
		if professionalLikesSkill(professionalID, skillName, session) {
			return dislikeSkill(professionalID, skillName, session)
		}
		return []byte("{\"Error\":\"The professional doesn't like this skill\"}")
	}
	return []byte("{\"Error\":\"Skill doesn't exists\"}")
}

func handleLikeSkill(professionalID *int, skillName *string, session neo4j.Session) []byte {
	if AccountCore.CheckIfSkillExists(skillName, session) {
		if !professionalLikesSkill(professionalID, skillName, session) {
			return likeSkill(professionalID, skillName, session)
		}
		return []byte("{\"Error\":\"The professional already likes this skill\"}")
	}
	return []byte("{\"Error\":\"Skill doesn't exists\"}")
}

func LikeSkillBackend(professionalID *int, skillName *string, like *bool) []byte {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		if *like {
			return handleLikeSkill(professionalID, skillName, session)
		}
		return handleDislikeSkill(professionalID, skillName, session)
	}
	log.Print(connectionError)
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
