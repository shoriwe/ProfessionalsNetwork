package Professional

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ProfessionalCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func KnownSkill(writer http.ResponseWriter, request *http.Request) {
	skillForm := new(RequestsForms.KnowSkillForm)
	if professionalID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, skillForm, true, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		if !skillForm.Know {
			ProfessionalCore.LikeSkillBackend(&professionalID, &skillForm.Skill, &skillForm.Know)
		}
		_, _ = writer.Write(ProfessionalCore.KnownSkillBackend(&professionalID, &skillForm.Skill, &skillForm.Know))
	}
}

func LikeSkill(writer http.ResponseWriter, request *http.Request) {
	skillForm := new(RequestsForms.LikeSkillForm)
	if professionalID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, skillForm, true, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		if skillForm.Like {
			ProfessionalCore.KnownSkillBackend(&professionalID, &skillForm.Skill, &skillForm.Like)
		}
		_, _ = writer.Write(ProfessionalCore.LikeSkillBackend(&professionalID, &skillForm.Skill, &skillForm.Like))
	}
}
