package Professional

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ProfessionalCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func ChangeGender(writer http.ResponseWriter, request *http.Request) {
	genderForm := new(RequestsForms.ChangeGenderForm)
	if professionalID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, genderForm, true, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		if InputHandler.IsGenderValid(&genderForm.Gender) {
			ProfessionalCore.ChangeGenderBackend(&professionalID, &genderForm.Gender)
			_, _ = writer.Write([]byte("{\"Success\":\"The gender was changed successfully\"}"))
		} else {
			_, _ = writer.Write([]byte("{\"Error\":\"The gender field is invalid\"}"))
		}
	}
}
