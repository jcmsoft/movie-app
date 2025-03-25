package handler

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var config = mysql.Config{
	User:                 "root",
	Passwd:               "root",
	Net:                  "tcp",
	Addr:                 "localhost:3306",
	DBName:               "movies",
	AllowNativePasswords: true,
}

func GetDatabaseConnection() (db *sql.DB) {
	var err error
	// Open the database connection
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("An error occured while opening the database connection %s", err)
	}
	// test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("An error occured while testing the database connection %s", err)
	} else {
		log.Println("Database connection established successfully")
	}
	return db
}

func CloseDBConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("An error occured while closing the database connection %s", err)
	}
}
