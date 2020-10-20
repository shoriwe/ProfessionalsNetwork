package Professionals

const (
	// ProfessionalSpeaksLanguageQuery = "MATCH (professional:Professional {id:$ProfessionalID})-[association:SPEAKS]->(language:Language {name:$Language}) RETURN professional"
	ProfessionalLikesSkillQuery = "MATCH (professional:Professional {id:$ProfessionalID})-[association:LIKE]->(skill:Skill {name:$Skill}) RETURN professional"
	ProfessionalKnowsSkillQuery = "MATCH (professional:Professional {id:$ProfessionalID})-[association:KNOW]->(skill:Skill {name:$Skill}) RETURN professional"
)
