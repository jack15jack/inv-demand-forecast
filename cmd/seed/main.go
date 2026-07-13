package main

import (
	"log"

	"github.com/jack15jack/inv-demand-forecast/internal/db"
)

func main() {

	log.Println("Starting seed process...")

	database := db.Connect()

	generator := NewGenerator(database)

	if err := generator.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("Seed complete.")
}
