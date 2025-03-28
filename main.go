package main

import (
	"github.com/gofiber/fiber/v3"
	
)

func main() {

	app := fiber.New()

	app.Get("/healthcheck", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})
}
