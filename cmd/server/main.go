package main

import (
	"log"
	"os"

	"github.com/jack15jack/inv-demand-forecast/internal/config"
	"github.com/jack15jack/inv-demand-forecast/internal/db"
	"github.com/jack15jack/inv-demand-forecast/internal/router"
)

func main() {

	config.LoadEnv()

	database := db.Connect()

	r := router.SetupRouter(database)

	port := os.Getenv("SERVER_PORT")

	log.Printf("Server listening on %s", port)

	r.Run(":" + port)
}
