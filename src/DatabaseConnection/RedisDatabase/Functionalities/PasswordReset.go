package Functionalities

import (
	"github.com/go-redis/redis/v8"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase"
	"github.com/shoriwe/ProNet/src/InputHandler/Templates"
	"log"
	"time"
)

func UserResetExists(passwordChangeKey *string) bool {
	connection := RedisDatabase.GetResetPasswordKeysConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.KeyExists(connection, passwordChangeKey).Val() == 1
}

func SetUserResetInformation(resetKey *string, resetInformation *Templates.UserResetInformation) {
	connection := RedisDatabase.GetResetPasswordKeysConnection()
	defer RedisDatabase.CloseConnection(connection)
	result := connection.Set(RedisDatabase.ConnectionContext, *resetKey, resetInformation, time.Hour)
	if result.Err() != nil {
		log.Print(result.Err())
	}
}

func GetUserResetInformation(passwordChangeKey *string) *redis.StringCmd {
	connection := RedisDatabase.GetResetPasswordKeysConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.GetValueByKey(connection, passwordChangeKey)
}

func DeleteUserResetInformation(passwordResetKey *string) {
	connection := RedisDatabase.GetResetPasswordKeysConnection()
	defer RedisDatabase.CloseConnection(connection)
	RedisDatabase.DeleteKey(connection, passwordResetKey)
}
