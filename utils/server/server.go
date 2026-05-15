package server

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 10,
	})
	return app
}

func RunHttpApp(lc fx.Lifecycle, app *fiber.App, logger *zap.Logger) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Path not found")
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Starting Fiber HTTP server", zap.String("addr", ":8080"))
				if err := app.Listen(":8080"); err != nil {
					logger.Fatal("Failed to start Fiber server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down Fiber HTTP server")
			return app.Shutdown()
		},
	})
}
