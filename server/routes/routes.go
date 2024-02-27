package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ordomigato/parking-app/middleware"
	"github.com/ordomigato/parking-app/utils"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// auth
	api.Post("/register", RegisterClient)
	api.Post("/login", LoginClient)
	api.Get("/logout", middleware.DeserializeClient, LogoutClient)
	api.Get("/status", middleware.DeserializeClient, ClientStatus)
	api.Post("/email/verify", VerifyClient)

	// workspace
	api.Post("/workspace", middleware.DeserializeClient, CreateWorkspace)
	api.Get("/workspace", middleware.DeserializeClient, GetWorkspaces)
	// TODO add middleware to determine ownership of workspace
	api.Put("/workspace/:wsid", middleware.DeserializeClient, middleware.WorkspaceAdmin, UpdateWorkspace)
	api.Delete("/workspace/:wsid", middleware.DeserializeClient, middleware.WorkspaceAdmin, DeleteWorkspace)

	// form
	api.Post("/workspace/:wsid/form", middleware.DeserializeClient, middleware.WorkspaceAdmin, CreateForm)
	api.Get("/workspace/:wsid/form", middleware.DeserializeClient, middleware.WorkspaceAdmin, GetForms)
	api.Get("/workspace/:wsid/form/:formid", middleware.DeserializeClient, middleware.WorkspaceAdmin, GetForm)
	api.Put("/workspace/:wsid/form/:formid", middleware.DeserializeClient, middleware.WorkspaceAdmin, UpdateForm)
	api.Delete("/workspace/:wsid/form/:formid", middleware.DeserializeClient, middleware.WorkspaceAdmin, DeleteForm)
	api.Put("/workspace/:wsid/form/:formid/path", middleware.DeserializeClient, middleware.WorkspaceAdmin, UpdateFormPath)

	// permit
	api.Get("/workspace/:wsid/form/:formid/permit", middleware.DeserializeClient, middleware.WorkspaceAdmin, GetPermits)
	api.Get("workspace/:wsid/form/:formid/permit/:permitid", middleware.DeserializeClient, middleware.WorkspaceAdmin, GetPermit)

	api.Get("/workspace/:wsid/form/:formid/permit/download", middleware.DeserializeClient, middleware.WorkspaceAdmin, DownloadPermits)
	api.Delete("/workspace/:wsid/form/:formid/permit/:permitid", middleware.DeserializeClient, middleware.WorkspaceAdmin, DeletePermit)

	// submit permit
	api.Post("/form/:formid/permit", CreatePermit)

	// get form info
	api.Get("/form/:workspacePath/:formPath", GetFormInfo)

	// blacklist
	api.Get("/workspace/:wsid/form/:formid/blacklist", middleware.DeserializeClient, middleware.WorkspaceAdmin, GetBlacklist)
	api.Post("/workspace/:wsid/form/:formid/blacklist", middleware.DeserializeClient, middleware.WorkspaceAdmin, CreateBlacklistEntry)
	api.Delete("/workspace/:wsid/form/:formid/blacklist/:plate", middleware.DeserializeClient, middleware.WorkspaceAdmin, DeleteBlacklistEntry)

	api.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("Path: %v does not exists on this server", path)))
	})
}
