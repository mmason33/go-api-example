package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	treatments "github.com/weareenvoy/nuface-recommendation-server-golang/handlers"
	"github.com/weareenvoy/nuface-recommendation-server-golang/middleware"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New(), middleware.AuthReq())

	// routes
	api.Get("/", treatments.GetAllTreatments)
	api.Get("/:id", treatments.GetSingleTreatment)
	api.Post("/", treatments.CreateTreatment)
	api.Delete("/:id", treatments.DeleteTreatment)
}
