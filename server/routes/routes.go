package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ordomigato/parking-app/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register", RegisterClient)
	api.Post("/login", LoginClient)
	api.Get("/logout", middleware.DeserializeClient, LogoutClient)
	api.Get("/status", middleware.DeserializeClient, ClientStatus)
	api.Post("/workspace", middleware.DeserializeClient, CreateWorkspace)
	// api.Get("/workspace", middleware.DeserializeClient, getWorkspaces)

	api.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})
}
