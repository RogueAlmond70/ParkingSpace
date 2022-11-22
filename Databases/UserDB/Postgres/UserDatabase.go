package Postgres

import (
	"database/sql"
	"fmt"
	"home/aaron/snap/go/GolandProjects/ParkingSpace/Constants"
	"log"
)

func getConn() string {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		Constants.DBHost,
		Constants.DBPort,
		Constants.DBUser,
		Constants.DBPass,
		Constants.DBName,
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
