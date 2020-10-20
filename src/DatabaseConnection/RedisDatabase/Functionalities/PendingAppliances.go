package Functionalities

import (
	"github.com/go-redis/redis/v8"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"log"
	"time"
)

func GetPendingAppliance(applicationKey *string) *redis.StringCmd {
	connection := RedisDatabase.GetPendingApplianceConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.GetValueByKey(connection, applicationKey)
}

func IsPendingAppliance(applicationKey *string) bool {
	connection := RedisDatabase.GetPendingApplianceConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.KeyExists(connection, applicationKey).Val() == 1
}

func DeletePendingAppliance(applicationKey *string) {
	connection := RedisDatabase.GetPendingApplianceConnection()
	defer RedisDatabase.CloseConnection(connection)
	RedisDatabase.DeleteKey(connection, applicationKey)
}

func SetPendingApplianceKeyValue(connection *redis.Client, key *string, value *Templates.UserAppliance, expiration time.Duration) {
	result := connection.Set(RedisDatabase.ConnectionContext, *key, value, expiration)
	if result.Err() != nil {
		log.Print(result.Err())
	}
}

func NewPendingAppliance(username *string, userInformation *Templates.UserInformation) {
	connection := RedisDatabase.GetPendingApplianceConnection()
	defer RedisDatabase.CloseConnection(connection)
	SetPendingRegistrationKeyValue(connection, username, userInformation, 24*time.Hour)
}

func UpdatePendingAppliance(username *string, userInformation *Templates.UserInformation, expiration time.Time) {
	connection := RedisDatabase.GetPendingApplianceConnection()
	defer RedisDatabase.CloseConnection(connection)
	SetPendingRegistrationKeyValue(connection, username, userInformation, time.Duration(expiration.Unix()))
}
