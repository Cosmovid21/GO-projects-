package main

import (
	"database/sql"
	"fmt"
	"log"
	//"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	// Capture connection properties.
	cfg := mysql.Config{
		//User:   os.Getenv("deepak"),
		//Passwd: os.Getenv("123456789"),
		User:   "deepak",
        Passwd: "123456789",
		Net:    "tcp",
		Addr:   "192.168.0.34:8000",
		DBName: "mines",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
