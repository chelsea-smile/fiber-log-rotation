package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	app := fiber.New()

	// Logger
	logger := CreateLogger()
	defer logger.Sync()

	// Middleware
	app.Use(func(c *fiber.Ctx) error {
		logger.Info("incoming request", zap.String(c.Path(), c.Method()))
		return c.Next()
	})

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello 🍏")
	})

	log.Fatal(app.Listen(":3000"))
}
