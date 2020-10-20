package AccountCore

import (
	"encoding/hex"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"golang.org/x/crypto/sha3"
	"log"
)

func UserIDExists(userID *int) (bool, string, int) {
	rows, connectionError := SQLDatabase.QuerySelectConnection(Queries.CheckIfUserIDExistsQuery, *userID)
	defer SQLDatabase.CloseConnection(rows)
	if connectionError == nil {
		if rows.Next() {
			username := new(string)
			accountType := new(int64)
			transformationError := rows.Scan(&username, &accountType)
			if transformationError == nil {
				return true, *username, int(*accountType)
			}
		}
	} else {
		log.Print(connectionError)
	}
	return false, "", -1
}
func UsernameExists(username *string) (bool, int, int) {
	rows, connectionError := SQLDatabase.QuerySelectConnection(Queries.CheckIfUsernameExistsAndGetIDAndAccountTypeQuery, *username)
	defer SQLDatabase.CloseConnection(rows)
	if connectionError == nil {
		if rows.Next() {
			accountID := new(int64)
			accountType := new(int64)
			transformationError := rows.Scan(&accountID, &accountType)
			if transformationError == nil {
				return true, int(*accountID), int(*accountType)
			}
		}
	} else {
		log.Print(connectionError)
	}
	return false, -1, -1
}

func UsernameAndPasswordExists(username *string, password *string) bool {
	hashHandler := sha3.New256()
	hashHandler.Write([]byte(*password))
	passwordHash := hex.EncodeToString(hashHandler.Sum(nil))
	rows, connectionError := SQLDatabase.QuerySelectConnection(Queries.CheckIfUsernamePasswordAreValidQuery, *username, passwordHash)
	defer SQLDatabase.CloseConnection(rows)
	if connectionError == nil {
		return rows.Next()
	} else {
		log.Print(connectionError)
	}
	return false
}
