package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/likhithkp/clip/utils/jwt"
	"github.com/likhithkp/clip/utils/other"
)

func AuthMiddleware(jwtManager *jwt.GenerateJwtTokenManager, utils *other.ResponseStruct) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("auth_token")

		if token == "" {
			authHeader := c.Get("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if token == "" {
			return utils.Response(c, fiber.StatusUnauthorized, false, "Missing auth token", nil)
		}

		claims, err := jwtManager.ValidateJWT(token)
		if err != nil {
			return utils.Response(c, fiber.StatusUnauthorized, false, "Invalid or expired token", nil)
		}

		c.Locals("userID", claims["sub"])
		c.Locals("email", claims["email"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}

func NewAuthMiddleware(jwtManager *jwt.GenerateJwtTokenManager, utils *other.ResponseStruct) fiber.Handler {
	return AuthMiddleware(jwtManager, utils)
}
