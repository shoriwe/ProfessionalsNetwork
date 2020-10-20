package Professional

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ProfessionalCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func ChangeNationality(writer http.ResponseWriter, request *http.Request) {
	nationalityForm := new(RequestsForms.ChangeNationalityForm)
	if professionalID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, nationalityForm, true, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		if InputHandler.IsNationalityValid(&nationalityForm.Nationality) {
			_, _ = writer.Write(ProfessionalCore.ChangeNationalityBackend(&professionalID, &nationalityForm.Nationality))
		} else {
			_, _ = writer.Write([]byte("{\"Error\": \"Nationality field is not valid\"}"))
		}
	}

}
