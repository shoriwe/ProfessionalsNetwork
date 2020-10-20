package Functionalities

import (
	"github.com/go-redis/redis/v8"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"log"
	"time"
)

func GetPendingRegistration(usernameHash *string) *redis.StringCmd {
	connection := RedisDatabase.GetPendingRegistrationsConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.GetValueByKey(connection, usernameHash)
}

func IsPendingRegistration(usernameHash *string) bool {
	connection := RedisDatabase.GetPendingRegistrationsConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.KeyExists(connection, usernameHash).Val() == 1
}

func DeletePendingRegistration(usernameHash *string) {
	connection := RedisDatabase.GetPendingRegistrationsConnection()
	defer RedisDatabase.CloseConnection(connection)
	RedisDatabase.DeleteKey(connection, usernameHash)
}

func SetPendingRegistrationKeyValue(connection *redis.Client, key *string, value *Templates.UserInformation, expiration time.Duration) {
	result := connection.Set(RedisDatabase.ConnectionContext, *key, value, expiration)
	if result.Err() != nil {
		log.Print(result.Err())
	}
}

func NewPendingRegistration(username *string, userInformation *Templates.UserInformation) {
	connection := RedisDatabase.GetPendingRegistrationsConnection()
	defer RedisDatabase.CloseConnection(connection)

	SetPendingRegistrationKeyValue(connection, username, userInformation, 24*time.Hour)
}

func UpdatePendingRegistration(username *string, userInformation *Templates.UserInformation, expiration time.Time) {
	connection := RedisDatabase.GetPendingRegistrationsConnection()
	defer RedisDatabase.CloseConnection(connection)
	SetPendingRegistrationKeyValue(connection, username, userInformation, time.Duration(expiration.Unix()))
}
