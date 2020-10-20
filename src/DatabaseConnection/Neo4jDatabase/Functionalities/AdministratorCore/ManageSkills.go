package AdministratorCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/General"
)

func createSkillNodeFilter(_ *neo4j.Result) interface{} {
	return true
}

func handleCreateSkillNode(skillName *string, session neo4j.Session) bool {
	data := map[string]interface{}{
		"Skill": *skillName,
	}
	if !AccountCore.CheckIfSkillExists(skillName, session) {
		result := Neo4jDatabase.ExecuteQuery(session, General.CreateSkillNodeQuery, data, createSkillNodeFilter)
		if result != nil {
			return result.(bool)
		}
	}
	return false
}

func CreateSkillNode(skillName *string) bool {
	session, connectionError := Neo4jDatabase.Database.NewSession(neo4j.SessionConfig{})
	defer Neo4jDatabase.CloseSession(session)
	if connectionError == nil {
		return handleCreateSkillNode(skillName, session)
	}
	return false
}
