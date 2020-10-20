package Contractor

import (
	"github.com/shoriwe/ProNet/src/API/SessionHandling"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase/Functionalities/ContractorCore"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates/RequestsForms"
	"net/http"
)

func FindProfessionals(writer http.ResponseWriter, request *http.Request) {
	searchFilter := new(RequestsForms.SearchFilterForm)
	if _, _, _, isAccountTypeCorrect, isRequestValid := SessionHandling.RequestHandler(writer, request, searchFilter, true, SQLDatabase.ContractorAccount); isAccountTypeCorrect && isRequestValid {
		message, responseStatusCode := ContractorCore.FindProfessionalsBackend(searchFilter)
		writer.WriteHeader(responseStatusCode)
		_, _ = writer.Write(message)
	}
}
