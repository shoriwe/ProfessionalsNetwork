package Contractors

const (
	CreateContractorNodeQuery = "CREATE (n:Contractor {id: $ContractorID, ownerID: $ContractorID, name: $ContractorName})"
	CreateTeamNodeQuery       = "CREATE (n:Team {id: $TeamID, ownerID: $ContractorID, name: $TeamName})"
)
