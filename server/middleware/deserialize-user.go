package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/ordomigato/parking-app/initializers"
	"github.com/ordomigato/parking-app/models"
	"github.com/ordomigato/parking-app/utils"
)

func DeserializeClient(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(
			utils.GenerateServerErrorResponse("You are not logged in"))
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			utils.GenerateServerErrorResponse(fmt.Sprintf("invalidate token: %v", err)))
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(
			utils.GenerateServerErrorResponse("invalid token claim"))

	}

	var client models.Client
	initializers.DB.First(&client, "client_id = ?", fmt.Sprint(claims["sub"]))

	if client.ClientID.String() != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(
			utils.GenerateServerErrorResponse("the user belonging to this token no logger exists"))
	}

	c.Locals("client", models.FilterClientRecord(&client))

	return c.Next()
}
