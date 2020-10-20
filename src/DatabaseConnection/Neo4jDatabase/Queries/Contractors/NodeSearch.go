package Contractors

const (
	LanguageFilter                         = "MATCH (professional:Professional)-->(language:Language) WITH CASE WHEN SIZE($Languages) = 0  THEN [professional, 0] ELSE CASE WHEN (language).name IN $Languages THEN [professional, COUNT(language)] END END AS result WHERE result[1] >= SIZE($Languages) WITH result[0] AS professional "
	LocationFilter                         = "MATCH (professional) WITH CASE WHEN $Locations = [] THEN professional ELSE CASE WHEN (professional).location IN $Locations THEN professional END END AS professional "
	NationalityFilter                      = "MATCH (professional) WITH CASE WHEN $Nationalities = [] THEN professional ELSE CASE WHEN (professional).nationality IN $Nationalities THEN professional END END AS professional "
	AvailableFilter                        = "MATCH (professional) WITH CASE WHEN $Available IS NULL THEN professional ELSE CASE WHEN (professional).available = $Available THEN professional END END AS professional "
	GenderFilter                           = "MATCH (professional) WITH CASE WHEN $Gender IS NULL THEN professional ELSE CASE WHEN (professional).gender = $Gender THEN professional END END AS professional "
	RemoteFilter                           = "MATCH (professional) WITH CASE WHEN $Remote IS NULL THEN professional ELSE CASE WHEN (professional).remote = $Remote THEN professional END END AS professional "
	SkillIncludeNotLikedFilter             = "MATCH (professional {liked_only:false})-->(skill:Skill) WITH CASE WHEN SIZE($Skills) = 0  THEN [professional, 0] ELSE CASE WHEN (skill).name IN $Skills THEN [professional, COUNT(skill)] END END AS result WHERE result[1] >= SIZE($Skills) WITH result[0] AS professional "
	SkillLikedOnlyFilter                   = "MATCH (professional)-->(skill:Skill) WHERE ( (professional).liked_only = true AND (professional)-[:LIKE]->(skill:Skill) ) OR ( (professional).liked_only = false AND (professional)-->(skill:Skill) ) WITH CASE WHEN SIZE($Skills) = 0  THEN [professional, 0] ELSE CASE WHEN (skill).name IN $Skills THEN [professional, COUNT(skill)] END END AS result WHERE result[1] >= SIZE($Skills) WITH result[0] AS professional "
	CollectResults                         = "MATCH (professional)-[:KNOW]->(skill:Skill) WITH professional,COUNT(skill) AS numberOfSkills WITH COLLECT([professional, numberOfSkills]) AS results WITH results[(($Page - 1) * 25)..($Page * 25)] AS results, ((SIZE(results)/25) +  1) AS numberOfResults RETURN results, numberOfResults"
	ProfessionalSearchIncludeNotLikedQuery = LanguageFilter + LocationFilter + NationalityFilter + AvailableFilter + GenderFilter + RemoteFilter + SkillIncludeNotLikedFilter + CollectResults
	ProfessionalSearchLikedOnlyQuery       = LanguageFilter + LocationFilter + NationalityFilter + AvailableFilter + GenderFilter + RemoteFilter + SkillLikedOnlyFilter + CollectResults
)

const (
	CheckProfessionalIsInTeamQuery = "MATCH (professional:Professional {id: $ProfessionalID, ownerID: $ProfessionalID})-[:WORK_IN]->(team:Team {id: $TeamID, name: $TeamName, ownerID: $ContractorID}) RETURN professional"
)

const (
	GetContractorNameQuery = "MATCH (contractor:Contractor {id: $ContractorID}) WITH contractor.name AS name_ RETURN name_"
)
