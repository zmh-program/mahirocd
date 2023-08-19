package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

var manager Manager

func main() {
	app := fiber.New()
	app.Use(recover.New())

	manager = NewManager()
	app.Post("/events", func(c *fiber.Ctx) error {
		var webhook GithubWebhook
		if err := c.BodyParser(&webhook); err != nil {
			return err
		}

		status := manager.RunAsync(webhook.Repository.Name)
		return c.JSON(&fiber.Map{"status": status})
	})

	log.Fatal(app.Listen(":3000"))
}
