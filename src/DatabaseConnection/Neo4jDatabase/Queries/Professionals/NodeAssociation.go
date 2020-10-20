package Professionals

const (
	SpeaksLanguageQuery      = "MATCH (professional:Professional {id:$ProfessionalID}),(language:Language {name:$Language}) CREATE (professional)-[:SPEAKS]->(language)"
	DoesntSpeakLanguageQuery = "MATCH (professional:Professional {id:$ProfessionalID})-[association:SPEAKS]->(language:Language {name:$Language}) DELETE association"
	LikeSkillQuery           = "MATCH (professional:Professional {id:$ProfessionalID}),(skill:Skill {name:$Skill}) CREATE (professional)-[:LIKE]->(skill)"
	KnownSkillQuery          = "MATCH (professional:Professional {id:$ProfessionalID}),(skill:Skill {name:$Skill}) CREATE (professional)-[:KNOW]->(skill)"
	DislikeSkillQuery        = "MATCH (professional:Professional {id:$ProfessionalID})-[association:LIKE]->(skill:Skill {name:$Skill}) DELETE association"
	UnknownSkillQuery        = "MATCH (professional:Professional {id:$ProfessionalID})-[association:KNOW]->(skill:Skill {name:$Skill}) DELETE association"
)

const (
	AddProfessionalToTeamQuery = "MATCH (professional:Professional {id:$ProfessionalID, ownerID:$ProfessionalID}) WITH professional MATCH (team:Team {id:$TeamID, ownerID:$ContractorID, name:$TeamName}) CREATE  (professional)-[:WORK_IN]->(team)"
)
