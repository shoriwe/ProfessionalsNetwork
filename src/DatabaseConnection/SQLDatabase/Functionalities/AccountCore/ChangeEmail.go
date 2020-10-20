package AccountCore

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"log"
)

func GetUserEmail(userID *int, username *string) (bool, string) {
	results, connectionError := SQLDatabase.QuerySelectConnection(Queries.GetUserEmailQuery, *userID, *username)
	if connectionError == nil {
		if results.Next() {
			email := new(string)
			parsingError := results.Scan(&email)
			if parsingError == nil {
				return true, *email
			} else {
				log.Print(parsingError)
			}
		}
	} else {
		log.Print(connectionError)
	}
	return false, ""
}

func UpdateEmail(userID *int, username *string, newEmail *string) bool {
	_, connectionError := SQLDatabase.ExecuteUpdateQuery(Queries.UpdateEmailQuery, *newEmail, *userID, *username)
	if connectionError == nil {
		return true
	}
	log.Print(connectionError)
	return false
}
