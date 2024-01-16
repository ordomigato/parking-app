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
	api.Put("/workspace/:wsID", middleware.DeserializeClient, middleware.WorkspaceAdmin, UpdateWorkspace)
	api.Delete("/workspace/:wsID", middleware.DeserializeClient, middleware.WorkspaceAdmin, DeleteWorkspace)

	// form
	api.Post("/workspace/:wsID/form", middleware.DeserializeClient, middleware.WorkspaceAdmin, CreateForm)
	api.Get("/workspace/:wsID/form", middleware.DeserializeClient, middleware.WorkspaceAdmin, GetForms)
	api.Get("/workspace/:wsID/form/:formId", middleware.DeserializeClient, middleware.WorkspaceAdmin, GetForm)
	api.Put("/workspace/:wsID/form/:formId", middleware.DeserializeClient, middleware.WorkspaceAdmin, UpdateForm)
	api.Delete("/workspace/:wsID/form/:formId", middleware.DeserializeClient, middleware.WorkspaceAdmin, DeleteForm)
	api.Put("/workspace/:wsID/form/:formId/path", middleware.DeserializeClient, middleware.WorkspaceAdmin, UpdateFormPath)

	// permit
	api.Get("/workspace/:wsID/form/:formId/permit", middleware.DeserializeClient, middleware.WorkspaceAdmin, GetPermits)

	// submit permit
	api.Post("/form/:formId/permit", CreatePermit)

	// get form info
	api.Get("/form/:workspacePath/:formPath", GetFormInfo)

	api.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":        "fail",
			"error_message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})
}
