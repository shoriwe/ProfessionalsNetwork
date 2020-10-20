package AccountCore

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Queries/General"
)

func checkIfSkillExistsFilter(result *neo4j.Result) interface{} {
	return (*result).Next()
}

func CheckIfSkillExists(skillName *string, session neo4j.Session) bool {
	data := map[string]interface{}{
		"Skill": *skillName,
	}
	result := Neo4jDatabase.ExecuteQuery(session, General.CheckIfSkillExistsQuery, data, checkIfSkillExistsFilter)
	if result != nil {
		return result.(bool)
	}
	return false
}
