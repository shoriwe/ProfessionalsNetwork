package Functionalities

import "github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase"

func IsAPIKeyValid(apiKey *string) bool {
	connection := RedisDatabase.GetActiveAPIKeysConnection()
	defer RedisDatabase.CloseConnection(connection)
	return RedisDatabase.KeyExists(connection, apiKey).Val() == 1
}
