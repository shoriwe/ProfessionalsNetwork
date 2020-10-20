package SQLDatabase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var updateConnection *sql.DB
var insertConnection *sql.DB
var selectConnection *sql.DB

func ExecuteUpdateQuery(query string, args ...interface{}) (sql.Result, error) {
	return updateConnection.Exec(query, args...)
}

func ExecuteInsertQuery(query string, args ...interface{}) (sql.Result, error) {
	return insertConnection.Exec(query, args...)
}

func CloseConnection(connection *sql.Rows) {
	if connection != nil {
		_ = connection.Close()
	}
}

func QuerySelectConnection(query string, args ...interface{}) (*sql.Rows, error) {
	return selectConnection.Query(query, args...)
}

func Connect(
	databaseHost string,
	sqlUpdateUsername string, sqlUpdatePassword string,
	sqlInsertUsername string, sqlInsertPassword string,
	sqlSelectUsername string, sqlSelectPassword string,
	sqlDatabase string) {

	databaseUri := "@tcp(" + databaseHost + ")/" + sqlDatabase

	updateConnection, _ = sql.Open("mysql", sqlUpdateUsername+":"+sqlUpdatePassword+databaseUri)
	if connectionError := updateConnection.Ping(); connectionError != nil {
		log.Fatal(connectionError)
	}

	insertConnection, _ = sql.Open("mysql", sqlInsertUsername+":"+sqlInsertPassword+databaseUri)
	if connectionError := insertConnection.Ping(); connectionError != nil {
		log.Fatal(connectionError)
	}

	selectConnection, _ = sql.Open("mysql", sqlSelectUsername+":"+sqlSelectPassword+databaseUri)
	if connectionError := selectConnection.Ping(); connectionError != nil {
		log.Fatal(connectionError)
	}
}
