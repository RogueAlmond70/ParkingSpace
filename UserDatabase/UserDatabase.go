package UserDatabase

import (
	"database/sql"
	"fmt"
	"log"
)

func getConn() string {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		viper.GetString(config.DBHost),
		viper.GetInt64(config.DBPort),
		viper.GetString(config.DBUser),
		viper.GetString(config.DBPass),
		viper.GetString(config.DBName))
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

func CreateTable(db *sql.DB) {
	query := Constants.CreateTable
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating table", err)
	}
}
