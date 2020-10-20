package Professionals

const (
	GetProfessionalLanguagesQuery = "MATCH (professional:Professional {id: $ProfessionalID})-->(language:Language) WITH COLLECT((language).name) AS languages RETURN languages"
)
