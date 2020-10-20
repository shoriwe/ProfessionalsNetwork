package General

const (
	RemoveProfessionalFromTeamQuery = "MATCH (professional:Professional {id:$ProfessionalID})-[association:WORK_IN]->(team:Team {name:$TeamName, id: $TeamID, ownerID: $ContractorID}) DELETE association"
)
