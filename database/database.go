package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Db *sql.DB

func getDatabaseInfo() string {
	err := godotenv.Load(".env")

	if err != nil {
		return err.Error()
	}
	// Database Info
	host := os.Getenv("PSQL_HOST")
	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASSWORD")
	dbname := os.Getenv("PSQL_DB_NAME")
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	return dns
}

func init() {
	var err error
	Db, err = sql.Open("postgres", getDatabaseInfo())

	if err != nil {
		panic(err)
	}
}
