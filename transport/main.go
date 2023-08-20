package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/spf13/viper"
	"log"
)

var manager *BroadcastManager

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	manager = NewBroadcastManager()
	app := fiber.New()
	app.Use(recover.New())

	{
		app.Post("/events", EventHandler)
		app.Get("/connection", websocket.New(ConnectionHandler))
	}

	log.Fatal(app.Listen(":" + viper.GetString("port")))
}
