package main

import (
	"log"
	"os"

	"github.com/jack15jack/inv-demand-forecast/internal/config"
	"github.com/jack15jack/inv-demand-forecast/internal/db"
	"github.com/jack15jack/inv-demand-forecast/internal/inventory"
	"github.com/jack15jack/inv-demand-forecast/internal/router"
)

func main() {

	config.LoadEnv()

	database := db.Connect()

	if err := database.AutoMigrate(
		&inventory.Item{},
	); err != nil {
		log.Fatal(err)
	}

	r := router.SetupRouter(database)

	port := os.Getenv("SERVER_PORT")

	log.Printf("Server listening on %s", port)

	r.Run(":" + port)
}
