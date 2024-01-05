package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/routes"
)

func init() {
	godotenv.Load(".env")

	// DB SETUP
	config := &initializers.Config{
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUserPassword: os.Getenv("DB_PASS"),
		DBUserName:     os.Getenv("DB_USER"),
		// DBSSLMode:  os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
	}

	initializers.ConnectDB(config)
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}))
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
