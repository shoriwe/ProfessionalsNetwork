package AccountCore

import (
	"encoding/hex"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase/Queries"
	"golang.org/x/crypto/sha3"
	"log"
)

func UpdatePassword(userID *int, username *string, newPassword *string) (bool, []byte) {
	hashHandler := sha3.New256()
	hashHandler.Write([]byte(*newPassword))
	passwordHash := hex.EncodeToString(hashHandler.Sum(nil))
	_, connectionError := SQLDatabase.ExecuteUpdateQuery(Queries.UpdatePasswordQuery, passwordHash, *userID, *username)
	if connectionError == nil {
		return true, []byte("{\"Success\":\"Password updated successfully\"}")
	}
	log.Print(connectionError)
	return false, []byte("{\"Error\":\"Something goes wrong\"}")
}
