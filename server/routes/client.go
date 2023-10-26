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
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/utils"
)

func (r *Repository) RegisterClient(c *fiber.Ctx) error {
	var payload *models.ClientRegisterRequest

	err := c.BodyParser(&payload)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	if payload.Password != payload.PasswordConfirm {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "passwords do not match"})
		return err
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": err.Error()})
		return err
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

	err = r.DB.Create(&newClient).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not register account"})
		return err
	}

	return c.JSON(newClient)
}

func (r *Repository) LoginClient(c *fiber.Ctx) error {
	loginRequest := new(models.ClientLoginRequest)
	// Extract the credentials from the request body
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Find the user by credentials
	client, err := r.FindByCredentials(loginRequest.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}

	result := utils.CheckPasswordHash(loginRequest.Password, client.Password)
	if !result {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid email or Password",
		})
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
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": fmt.Sprintf("generating JWT Token failed: %v", err)})
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
		"user":  client,
		"token": tokenString,
	})
}

func (r *Repository) clientStatus(c *fiber.Ctx) error {
	client := c.Locals("client").(models.ClientResponse)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"client": client}})
}

func (r *Repository) LogoutClient(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func (r *Repository) FindByCredentials(username string) (*models.Client, error) {
	user := models.Client{Username: username}
	result := r.DB.Where("username = ?", strings.ToLower(username)).Find(&user)
	fmt.Print(result)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
