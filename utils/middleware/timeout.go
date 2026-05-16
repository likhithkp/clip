package middleware

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TimeoutMiddleware(timeout time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), timeout)
		defer cancel()

		c.SetUserContext(ctx)

		err := c.Next()

		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
				"error": "Request timeout",
			})
		}

		return err
	}
}
