package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func StartDB() {
	// Connect to Database
	fmt.Println("Connecting to database...")
	var err error
	DB, err = sql.Open("postgres", dataSource())
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
}

func dataSource() string {
	host := "localhost"
	pass := "pass"
	return "postgresql://" + host + ":5433/csv" +
		"?user=csv&sslmode=disable&password=" + pass
}
