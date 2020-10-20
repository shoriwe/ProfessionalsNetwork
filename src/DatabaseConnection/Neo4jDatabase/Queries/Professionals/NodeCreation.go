package Professionals

const (
	CreateProfessionalNodeQuery = "CREATE (n:Professional {id: $ProfessionalID, ownerID: $ProfessionalID, name: $ProfessionalName})"
)
