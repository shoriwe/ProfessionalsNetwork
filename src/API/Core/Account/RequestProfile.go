package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func RequestProfile(writer http.ResponseWriter, request *http.Request) {
	requestProfileForm := new(RequestsForms.RequestProfileForm)
	if accountID, _, accountType, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, requestProfileForm, true, SQLDatabase.ContractorAccount, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(AccountCore.RequestProfileBackend(&accountID, accountType))
	}
}
