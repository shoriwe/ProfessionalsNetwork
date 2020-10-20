package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func ChangeLocation(writer http.ResponseWriter, request *http.Request) {
	changeLocationForm := new(RequestsForms.ChangeLocationForm)
	if accountID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, changeLocationForm, true, SQLDatabase.ContractorAccount, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(AccountCore.ChangeLocationBackend(&accountID, &changeLocationForm.Location))
	}
}
