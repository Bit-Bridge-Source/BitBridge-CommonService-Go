package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func FiberJWTAuthenticator(secret []byte) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")[7:]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid authorization token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid authorization token",
			})
		}

		c.Locals("user_id", claims["user_id"])
		return c.Next()
	}
}
