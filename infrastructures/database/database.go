package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type SqlHandler struct {
	Conn *sql.DB
}

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

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("postgres", getDatabaseInfo())

	if err != nil {
		panic(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
