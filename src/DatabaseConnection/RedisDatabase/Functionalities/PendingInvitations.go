package Functionalities

import (
	"github.com/go-redis/redis/v8"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"log"
	"time"
)

func GetPendingInvitation(invitationKey *string) *redis.StringCmd {
	connection := RedisDatabase.GetPendingInvitationConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.GetValueByKey(connection, invitationKey)
}

func IsPendingInvitation(invitationKey *string) bool {
	connection := RedisDatabase.GetPendingInvitationConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.KeyExists(connection, invitationKey).Val() == 1
}

func DeletePendingInvitation(invitationKey *string) {
	connection := RedisDatabase.GetPendingInvitationConnection()
	defer RedisDatabase.CloseConnection(connection)
	RedisDatabase.DeleteKey(connection, invitationKey)
}

func SetPendingInvitationKeyValue(connection *redis.Client, key *string, value *Templates.UserInvitation, expiration time.Duration) {
	result := connection.Set(RedisDatabase.ConnectionContext, *key, value, expiration)
	if result.Err() != nil {
		log.Print(result.Err())
	}
}

func NewPendingInvitation(username *string, userInformation *Templates.UserInformation) {
	connection := RedisDatabase.GetPendingInvitationConnection()
	defer RedisDatabase.CloseConnection(connection)
	SetPendingRegistrationKeyValue(connection, username, userInformation, 24*time.Hour)
}

func UpdatePendingInvitation(username *string, userInformation *Templates.UserInformation, expiration time.Time) {
	connection := RedisDatabase.GetPendingInvitationConnection()
	defer RedisDatabase.CloseConnection(connection)
	SetPendingRegistrationKeyValue(connection, username, userInformation, time.Duration(expiration.Unix()))
}
