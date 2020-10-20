package Account

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/AccountCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func SearchSkills(writer http.ResponseWriter, request *http.Request) {
	searchForm := new(RequestsForms.SearchSkillsForm)
	if accountID, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, searchForm, true, SQLDatabase.ProfessionalAccount, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		_, _ = writer.Write(AccountCore.SearchSkillsBackend(&accountID, &searchForm.ObjectID, &searchForm.SearchQuery))
	}
}
