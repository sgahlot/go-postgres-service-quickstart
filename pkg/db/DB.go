package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/myeung18/service-binding-client/pkg/binding/convert"
	"log"
)

const (
	DRIVER_NAME     = "postgres"
	DEFAULT_DB_NAME = "test"
	DEFAULT_DB_URL  = "host=localhost port=5432 user=postgres password=password dbname=test sslmode=disable"
)

var (
	postgresSqlClient *sql.DB
)

func getPostgresSqlConnectionStringForNonBindingsRun() string {
	return GetEnvOrDefault(DB_URL_KEY, DEFAULT_DB_URL)
}

func getPostgresSqlConnectionStringForBindingsRun() string {
	sqlConnectionStr, err := convert.GetPostgreSQLConnectionString()
	CheckErrorWithPanic(err, "while trying to get PostgresSQL connection string from Bindings")

	return sqlConnectionStr
}

func getPostgresSqlConnectionString() string {
	bindingsDir := GetEnvOrDefault("SERVICE_BINDING_ROOT", "")
	var sqlConnectionStr string
	if bindingsDir == "" {
		sqlConnectionStr = getPostgresSqlConnectionStringForNonBindingsRun()
	} else {
		log.Printf("System property for Bindings dir [%s] found", bindingsDir)
		sqlConnectionStr = getPostgresSqlConnectionStringForBindingsRun()
	}

	log.Printf("DB Connection String = %s", sqlConnectionStr)
	return sqlConnectionStr
}

func createPostgresSqlConnection() *sql.DB {
	db, err := sql.Open(DRIVER_NAME, getPostgresSqlConnectionString())
	CheckErrorWithPanic(err, "while connecting to PostgresSQL")

	db.SetMaxOpenConns(5)
	// defer db.Close()

	err = db.Ping()
	CheckErrorWithPanic(err, "while pinging PostgreSQL")

	return db
}

func checkAndRefreshConnection() {
	if postgresSqlClient == nil {
		// Try to get the connection as this could be the first time we're trying to connect
		postgresSqlClient = createPostgresSqlConnection()
	} else if err := postgresSqlClient.Ping(); err != nil {
		// Try to get the connection again as we might be disconnected
		postgresSqlClient = createPostgresSqlConnection()
	}

}

func GetPostgreSqlConnection() *sql.DB {
	checkAndRefreshConnection()

	return postgresSqlClient
}
