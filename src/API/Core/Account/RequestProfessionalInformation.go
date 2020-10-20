package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func RequestProfessionalInformation(writer http.ResponseWriter, request *http.Request) {
	professionalInformation := new(RequestsForms.AccountInformationForm)
	if _, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, professionalInformation, true, SQLDatabase.ContractorAccount, SQLDatabase.ProfessionalAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(AccountCore.RequestProfessionalInformationBackend(&professionalInformation.AccountID))
	}
}
