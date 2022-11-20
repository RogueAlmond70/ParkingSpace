package UserDatabase

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	DBHost = "db.host"
	DBPass = "db.password"
	DBPort = "db.port"
	DBUser = "db.user"
	DBName = "db.name"
)

func getConn() string {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBPass,
		DBName,
	)
	return psqlconn
}

func DBSetup() *sql.DB {
	connStr := getConn()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

/*
func CreateTable(db *sql.DB) {
	query := Constants.CreateTable
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating table", err)
	}
}

*/
