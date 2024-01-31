package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/ordomigato/parking-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort, config.TimeZone)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	err = DB.AutoMigrate(&models.Client{})
	if err != nil {
		log.Fatal("Client Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.Workspace{})
	if err != nil {
		log.Fatal("Workspace Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.Form{})
	if err != nil {
		log.Fatal("Form Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.Path{})
	if err != nil {
		log.Fatal("Path Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.Role{})
	if err != nil {
		log.Fatal("Role Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.Blacklist{})
	if err != nil {
		log.Fatal("Blacklist Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	// newRole := models.Role{
	// 	RoleName: "Admin",
	// }

	// err = DB.Create(&newRole).Error
	// if err != nil {
	// 	log.Fatal("Admin Role Creation Failed:  \n", err.Error())
	// 	os.Exit(1)
	// }

	err = DB.AutoMigrate(&models.ClientWorkspace{})
	if err != nil {
		log.Fatal("ClientWorkspace Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.Permit{})
	if err != nil {
		log.Fatal("Permit Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Connected Successfully to the Database")
}
