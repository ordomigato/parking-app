package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ordomigato/parking-app/middleware"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register", r.RegisterClient)
	api.Post("/login", r.LoginClient)
	api.Get("/logout", middleware.DeserializeClient, r.LogoutClient)
	api.Get("/status", middleware.DeserializeClient, r.clientStatus)

	api.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})
}
