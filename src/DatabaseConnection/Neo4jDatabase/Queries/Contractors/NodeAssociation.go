package Contractors

const (
	ContractorOwnerOfTeamQuery = "MATCH (contractor:Contractor {id: $ContractorID, ownerID: $ContractorID}) WITH contractor MATCH (team:Team {id: $TeamID, ownerID: $ContractorID}) CREATE (contractor)-[:OWNS]->(team)"
)

const (
	DissolveTeamQuery = "MATCH (team:Team {id:$TeamID, ownerID: $OwnerID, name:$TeamName}) DETACH DELETE team"
)
