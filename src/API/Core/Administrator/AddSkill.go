package Administrator

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AdministratorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func AddSkill(writer http.ResponseWriter, request *http.Request) {
	skillForm := new(RequestsForms.AddSkillForm)
	if _, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, &skillForm, true, SQLDatabase.AdministratorAccount); isAccountTypeCorrect && isRequestValid {
		if AdministratorCore.CreateSkillNode(&skillForm.Name) {
			_, _ = writer.Write([]byte("{\"Success\":\"Skill created success fully\"}"))
		} else {
			_, _ = writer.Write([]byte("{\"Error\":\"Looks like skill already exists\"}"))
		}
	}
}
