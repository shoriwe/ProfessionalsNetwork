package AccountCore

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"log"
)

func TeamAlreadyExists(contractorID *int, teamName *string) int {
	rows, connectionError := SQLDatabase.QuerySelectConnection(Queries.CheckIfTeamExistsQuery, *contractorID, *teamName)
	if connectionError == nil {
		if rows.Next() {
			teamID := new(int)
			scanError := rows.Scan(teamID)
			if scanError == nil {
				return *teamID
			} else {
				log.Print(scanError)
			}
		}
	} else {
		log.Print(connectionError)
	}
	return -1
}
