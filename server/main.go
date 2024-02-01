package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/routes"
)

func init() {
	config := initializers.GetEnvConfig()

	initializers.ConnectDB(config)
	// notification.SendEMail(config)
}

func main() {
	config := initializers.GetEnvConfig()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.FrontendURL,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}))
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
