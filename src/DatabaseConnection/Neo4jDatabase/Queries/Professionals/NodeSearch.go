package Professionals

const (
	SearchTeamsQuery = "MATCH (professional {ownerID: $ProfessionalID, id: $ProfessionalID}) WITH professional MATCH (team:Team) WHERE NOT (professional)-->(team) WITH team WHERE toLower(team.name) CONTAINS $SearchQuery WITH COLLECT(team) AS teams RETURN teams"
)

const (
	GetProfessionalNameQuery = "MATCH (professional:Professional {id: $ProfessionalID}) WITH professional.name AS name_ RETURN name_"
)
