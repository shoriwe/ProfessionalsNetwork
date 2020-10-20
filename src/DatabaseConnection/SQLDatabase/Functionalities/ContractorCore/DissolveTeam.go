package ContractorCore

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
)

func DissolveTeam(teamID *int, contractorID *int, teamName *string) bool {
	_, connectionError := SQLDatabase.ExecuteUpdateQuery(Queries.UpdateTeamStatusQuery, *teamID, *teamName, *contractorID)
	return connectionError == nil
}
