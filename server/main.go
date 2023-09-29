package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/routes"
	"github.com/ordomigato/parking-app/store"
)

func main() {
	godotenv.Load(".env")

	// DB SETUP
	config := &store.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := store.NewConnection(config)

	if err != nil {
		log.Fatalf("DB Unable to load: %s", err)
	}

	// MODEL MIGRATIONS
	err = models.MigrateClient(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := routes.Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
