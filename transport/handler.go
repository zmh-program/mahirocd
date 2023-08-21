package main

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func EventHandler(c *fiber.Ctx) error {
	if !VerifySignature(c) {
		return errors.New("invalid signature")
	}

	var webhook GithubWebhook
	if err := c.BodyParser(&webhook); err != nil {
		return err
	}
	fmt.Println("Received webhook from repository:", webhook.Repository.FullName)

	go manager.Send(webhook.ToBytes())
	return c.SendStatus(fiber.StatusOK)
}

func ConnectionHandler(c *websocket.Conn) {
	manager.AddConnection(c)
	defer manager.RemoveConnection(c)

	for {
		if _, _, err := c.ReadMessage(); err != nil {
			break
		}
	}
}
