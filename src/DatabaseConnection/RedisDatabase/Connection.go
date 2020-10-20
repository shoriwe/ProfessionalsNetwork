package RedisDatabase

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var ConnectionContext = context.Background()
var address string
var password string

func Connect(redisHost string, redisPassword string) {
	address = redisHost
	password = redisPassword
	redisConnectionOptions := redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	}
	connection := redis.NewClient(&redisConnectionOptions)
	pingResult := connection.Ping(ConnectionContext)
	if pingResult.Err() != nil {
		log.Fatal(pingResult.Err())
	}
	_ = connection.Close()
}

func CloseConnection(connection *redis.Client) {
	_ = connection.Close()
}

func GetPendingRegistrationsConnection() *redis.Client {
	redisConnectionOptions := redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	}
	return redis.NewClient(&redisConnectionOptions)
}

func GetActiveAPIKeysConnection() *redis.Client {
	redisConnectionOptions := redis.Options{
		Addr:     address,
		Password: password,
		DB:       1,
	}
	return redis.NewClient(&redisConnectionOptions)
}

func GetResetPasswordKeysConnection() *redis.Client {
	redisConnectionOptions := redis.Options{
		Addr:     address,
		Password: password,
		DB:       2,
	}
	return redis.NewClient(&redisConnectionOptions)
}

func GetPendingEmailPhoneNumberChangeConnection() *redis.Client {
	redisConnectionOptions := redis.Options{
		Addr:     address,
		Password: password,
		DB:       3,
	}
	return redis.NewClient(&redisConnectionOptions)
}

func GetPendingInvitationConnection() *redis.Client {
	redisConnectionOptions := redis.Options{
		Addr:     address,
		Password: password,
		DB:       4,
	}
	return redis.NewClient(&redisConnectionOptions)
}

func GetPendingApplianceConnection() *redis.Client {
	redisConnectionOptions := redis.Options{
		Addr:     address,
		Password: password,
		DB:       5,
	}
	return redis.NewClient(&redisConnectionOptions)
}

func GetValueByKey(connection *redis.Client, key *string) *redis.StringCmd {
	return connection.Get(ConnectionContext, *key)
}

func DeleteKey(connection *redis.Client, key *string) *redis.IntCmd {
	return connection.Del(ConnectionContext, *key)
}

func KeyExists(connection *redis.Client, key *string) *redis.IntCmd {
	return connection.Exists(ConnectionContext, *key)
}

