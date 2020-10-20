package Professional

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ProfessionalCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func ChangeAvailable(writer http.ResponseWriter, request *http.Request) {
	availableForm := new(RequestsForms.ChangeAvailableForm)
	if professionalID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, availableForm, true, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(ProfessionalCore.ChangeAvailableBackend(&professionalID, &availableForm.Available))
	}
}
