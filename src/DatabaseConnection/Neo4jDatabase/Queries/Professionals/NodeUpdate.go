package Professionals

const (
	ChangeProfessionalLikedOnlyQuery   = "MATCH (professional:Professional {id: $ProfessionalID}) SET (professional).liked_only = $LikedOnly RETURN professional"
	ChangeProfessionalGenderQuery      = "MATCH (professional:Professional {id: $ProfessionalID}) SET (professional).gender = $Gender RETURN professional"
	ChangeProfessionalNationalityQuery = "MATCH (professional:Professional {id: $ProfessionalID}) SET (professional).nationality = $Nationality RETURN professional"
	ChangeRemoteQuery                  = "MATCH (professional:Professional {id: $ProfessionalID}) SET (professional).remote = $Remote RETURN professional"
	ChangeAvailableQuery               = "MATCH (professional:Professional {id: $ProfessionalID}) SET (professional).available = $Available RETURN professional"
)
