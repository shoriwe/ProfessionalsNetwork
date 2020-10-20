package Account

import (
	"encoding/json"
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"log"
	"net/http"
)

func RequestAccountType(writer http.ResponseWriter, request *http.Request) {
	accountTypeForm := new(RequestsForms.RequestAccountTypeForm)
	if accountID, username, accountType, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, accountTypeForm, true, SQLDatabase.ProfessionalAccount, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		rawProfessionalInformation := map[string]interface{}{
			"AccountType": SQLDatabase.AccountTypesByNumber[accountType],
			"Username":    username,
			"AccountID":   accountID,
		}
		professionalInformation, marshalError := json.Marshal(rawProfessionalInformation)
		if marshalError == nil {
			_, _ = writer.Write(professionalInformation)
		} else {
			log.Print(marshalError)
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte("{}"))
		}
	}
}
