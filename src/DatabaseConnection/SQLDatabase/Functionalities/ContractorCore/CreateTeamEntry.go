package ContractorCore

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"log"
)

func CreateTeamEntry(contractorID *int, teamName *string) bool {
	_, connectionError := SQLDatabase.ExecuteInsertQuery(Queries.CreateTeamQuery, *contractorID, *teamName, true)
	if connectionError == nil {
		return true
	}
	log.Print(connectionError)
	return false
}
