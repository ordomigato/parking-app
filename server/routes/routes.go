package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ordomigato/parking-app/middleware"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	api := app.Group("/api")
	api.Post("/register", r.RegisterClient)
	api.Post("/login", r.LoginClient)
	api.Get("/logout", r.LogoutClient)
	api.Get("/status", middleware.DeserializeClient, r.clientStatus)
}
