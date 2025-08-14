package main

import (
	"log"
	"server/db"
)

func main() {

	_, err := db.NewDataBase()
	if err != nil {
		log.Fatalf("Could not initialaze database: %s", err)
		return
	}

}
