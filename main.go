package main

import (
	"fmt"
	"log"
	"mahirocd/runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type GithubWebhook struct {
	HookId int64 `json:"hook_id"`
	Hook   struct {
		Type   string   `json:"type"`
		Id     int64    `json:"id"`
		Events []string `json:"events"`
	} `json:"hook"`
	Repository struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
	}
}

func main() {
	runtime.NewRuntime("test", "/", []string{"echo", "hello"}).ProcessAsync()

	app := fiber.New()
	app.Use(recover.New())
	app.Post("/events", func(c *fiber.Ctx) error {
		var webhook GithubWebhook
		if err := c.BodyParser(&webhook); err != nil {
			return err
		}
		fmt.Println(webhook)
		return c.JSON(webhook)
	})

	log.Fatal(app.Listen(":3000"))
}
