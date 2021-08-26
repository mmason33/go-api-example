package main

import (
	"log"

	"github.com/gofiber/fiber/v2" // import the fiber package
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/weareenvoy/nuface-recommendation-server-golang/database"
	"github.com/weareenvoy/nuface-recommendation-server-golang/router"

	_ "github.com/lib/pq"
)

// entry point to our program
func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// call the New() method - used to instantiate a new Fiber App
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	router.SetupRoutes(app)

	// listen on port 3000
	app.Listen(":3000")

}
