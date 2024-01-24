package routes

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/utils"
)

func RegisterClient(c *fiber.Ctx) error {
	var payload *models.ClientRegisterRequest

	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("passwords do not match"))
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			utils.GenerateServerErrorResponse(""))
	}

	now := time.Now()

	newClient := models.Client{
		Username:  strings.ToLower(payload.Username),
		Password:  hashedPassword,
		Verified:  false,
		CreatedAt: now,
		UpdatedAt: now,
		LastLogin: now,
	}

	err = initializers.DB.Create(&newClient).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Account already created"))
	}

	return c.JSON(models.FilterClientRecord(&newClient))
}

func LoginClient(c *fiber.Ctx) error {
	loginRequest := new(models.ClientLoginRequest)
	// Extract the credentials from the request body
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse(err.Error()))
	}

	// Find the user by credentials
	client, err := FindByCredentials(loginRequest.Username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("invalid email or password"))
	}

	result := utils.CheckPasswordHash(loginRequest.Password, client.Password)
	if !result {
		return c.Status(fiber.StatusBadRequest).JSON(
			utils.GenerateServerErrorResponse("Invalid email or Password"))
	}

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = client.ClientID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 1).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("generating JWT Token failed: %v", err)))
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   60 * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.JSON(fiber.Map{
		"user":  models.FilterClientRecord(client),
		"token": tokenString,
	})
}

func ClientStatus(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func LogoutClient(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.SendStatus(http.StatusOK)
}

func FindByCredentials(username string) (*models.Client, error) {
	user := models.Client{Username: username}
	result := initializers.DB.First(&user, "username = ?", strings.ToLower(username)).Find(&user)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
