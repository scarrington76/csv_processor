package main

import (
	"log"

	"csv_processor/db"
	"csv_processor/server"
)

func main() {
	db.StartDB()
	defer db.DB.Close()

	a := server.New()
	if err := a.Serve(); err != nil {
		db.DB.Close()
		log.Fatal("error serving application: ", err)
	}
}
