package main

import (
	"online-store-golang/configuration"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configuration.New()
	configuration.NewDatabase(config)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "hello",
		})
	})
	app.Listen(":3000")
}
