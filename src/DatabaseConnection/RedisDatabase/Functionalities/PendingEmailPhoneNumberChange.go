package Functionalities

import (
	"github.com/go-redis/redis/v8"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"log"
	"time"
)

func DeletePendingChangeEmailPhoneNumber(usernameHash *string) {
	connection := RedisDatabase.GetPendingEmailPhoneNumberChangeConnection()
	defer RedisDatabase.CloseConnection(connection)
	RedisDatabase.DeleteKey(connection, usernameHash)
}

func UpdatePendingChangeEmailPhoneNumber(username *string, userInformation *Templates.ChangeEmailPhoneNumber, expiration time.Time) {
	connection := RedisDatabase.GetPendingEmailPhoneNumberChangeConnection()
	defer RedisDatabase.CloseConnection(connection)
	SetPendingEmailPhoneNumberChangeKeyValue(connection, username, userInformation, time.Duration(expiration.Unix()))
}

func GetPendingEmailPhoneNumberChange(usernameHash *string) *redis.StringCmd {
	connection := RedisDatabase.GetPendingEmailPhoneNumberChangeConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.GetValueByKey(connection, usernameHash)
}

func IsPendingEmailPhoneNumberChange(usernameHash *string) bool {
	connection := RedisDatabase.GetPendingEmailPhoneNumberChangeConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.KeyExists(connection, usernameHash).Val() == 1
}

func SetPendingEmailPhoneNumberChangeKeyValue(connection *redis.Client, key *string, value *Templates.ChangeEmailPhoneNumber, expiration time.Duration) {
	result := connection.Set(RedisDatabase.ConnectionContext, *key, value, expiration)
	if result.Err() != nil {
		log.Print(result.Err())
	}
}

func NewPendingChangeEmailPhoneNumber(usernameHash *string, changeEmailPhoneNumber *Templates.ChangeEmailPhoneNumber) {
	connection := RedisDatabase.GetPendingEmailPhoneNumberChangeConnection()
	defer RedisDatabase.CloseConnection(connection)

	SetPendingEmailPhoneNumberChangeKeyValue(connection, usernameHash, changeEmailPhoneNumber, time.Hour)
}
