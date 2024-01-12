package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ordomigato/parking-app/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// auth
	api.Post("/register", RegisterClient)
	api.Post("/login", LoginClient)
	api.Get("/logout", middleware.DeserializeClient, LogoutClient)
	api.Get("/status", middleware.DeserializeClient, ClientStatus)

	// workspace
	api.Post("/workspace", middleware.DeserializeClient, CreateWorkspace)
	api.Get("/workspace", middleware.DeserializeClient, GetWorkspaces)
	// TODO add middleware to determine ownership of workspace
	api.Put("/workspace/:id", middleware.DeserializeClient, UpdateWorkspace)
	api.Delete("/workspace/:id", middleware.DeserializeClient, DeleteWorkspace)

	// form
	api.Post("/workspace/:id/form", middleware.DeserializeClient, CreateForm)
	api.Get("/workspace/:id/form", middleware.DeserializeClient, GetForms)
	api.Get("/workspace/:id/form/:formId", GetForm)
	api.Put("/workspace/:id/form/:formId", UpdateForm)
	api.Delete("/workspace/:id/form/:formId", DeleteForm)

	api.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":        "fail",
			"error_message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})
}
