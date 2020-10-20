package main

import (
	"github.com/shoriwe/ProNet/src/API"
	"github.com/shoriwe/ProNet/src/ArgumentParser"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/Neo4jDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/RedisDatabase"
	"github.com/shoriwe/ProNet/src/DatabaseConnection/SQLDatabase"
	"github.com/shoriwe/ProNet/src/MessageSending"
)

func main() {
	variables := ArgumentParser.GetEnvironmentVariables()
	MessageSending.ConfigureMessagingAPIVariables(variables["MessagingAPIKey"], variables["MessagingAPIURI"], variables["FrontendURI"])

	RedisDatabase.Connect(variables["RedisHost"], variables["RedisPassword"])
	Neo4jDatabase.Connect(variables["Neo4jURI"], variables["Neo4jUsername"], variables["Neo4jPassword"])
	SQLDatabase.Connect(
		variables["SQLHost"],
		variables["SQLUpdateUsername"], variables["SQLUpdatePassword"],
		variables["SQLInsertUsername"], variables["SQLInsertPassword"],
		variables["SQLSelectUsername"], variables["SQLSelectPassword"],
		variables["SQLDatabase"])

	API.Serve(variables["BindAddress"])
}
