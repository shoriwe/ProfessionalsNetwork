package AccountCore

import (
	"database/sql"
	"encoding/json"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"log"
)

func RequestAccount(rows *sql.Rows) (bool, string, string) {
	if rows.Next() {
		email := new(string)
		phoneNumber := new(string)
		transformationError := rows.Scan(email, phoneNumber)
		if transformationError == nil {
			return true, *email, *phoneNumber
		}
	}
	return false, "", ""
}

func RequestAccountBackend(accountID *int, username *string) []byte {
	rows, connectionError := SQLDatabase.QuerySelectConnection(Queries.GetAccountEmailPhoneNumberQuery, *accountID, *username)
	defer SQLDatabase.CloseConnection(rows)
	if connectionError == nil {
		found, email, phoneNumber := RequestAccount(rows)
		if found {
			marshalResult, marshalError := json.Marshal(map[string]string{
				"Email":       email,
				"PhoneNumber": phoneNumber,
			})
			if marshalError == nil {
				return marshalResult
			} else {
				log.Print(marshalError)
			}
		}
	} else {
		log.Print(connectionError)
	}
	return []byte("{\"Error\":\"Something goes wrong\"}")
}
