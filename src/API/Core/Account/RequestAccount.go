package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func RequestAccount(writer http.ResponseWriter, request *http.Request) {
	requestAccountForm := new(RequestsForms.RequestAccountForm)
	if accountID, username, _, _, isRequestValid := SessionHandling.RequestHandler(writer, request, requestAccountForm, true, SQLDatabase.ContractorAccount, SQLDatabase.ProfessionalAccount); isRequestValid {
		_, _ = writer.Write(AccountCore.RequestAccountBackend(&accountID, &username))
	}
}
