package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"mahirocd/workflow"
)

var manager workflow.Manager

func main() {
	app := fiber.New()
	app.Use(recover.New())

	manager = workflow.NewManager()

}
