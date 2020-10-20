package AccountCore

import (
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"log"
)

func GetContactInformation(accountID *int) (bool, string, string) {
	rows, connectionError := SQLDatabase.QuerySelectConnection(Queries.GetAccountEmailPhoneNumberWithAccountIDQuery, *accountID)
	defer SQLDatabase.CloseConnection(rows)
	if connectionError == nil {
		return RequestAccount(rows)
	} else {
		log.Print(connectionError)
	}
	return false, "", ""
}
