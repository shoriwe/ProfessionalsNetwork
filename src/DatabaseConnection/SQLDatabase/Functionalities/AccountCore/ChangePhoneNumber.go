package AccountCore

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"log"
)

func GetUserPhoneNumber(userID *int, username *string) (bool, string) {
	results, connectionError := SQLDatabase.QuerySelectConnection(Queries.GetUserPhoneNumberQuery, *userID, *username)
	if connectionError == nil {
		if results.Next() {
			phoneNumber := new(string)
			parsingError := results.Scan(&phoneNumber)
			if parsingError == nil {
				return true, *phoneNumber
			} else {
				log.Print(parsingError)
			}
		}
	} else {
		log.Print(connectionError)
	}
	return false, ""
}

func UpdatePhoneNumber(userID *int, username *string, newPhoneNumber *string) bool {
	_, connectionError := SQLDatabase.ExecuteUpdateQuery(Queries.UpdatePhoneNumberQuery, *newPhoneNumber, *userID, *username)
	if connectionError == nil {
		return true
	}
	log.Print(connectionError)
	return false
}
