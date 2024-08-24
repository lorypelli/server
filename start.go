package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
)

func Start(dir, name string, port uint16) {
	app := fiber.New(fiber.Config{
		AppName:      name,
		ServerHeader: name,
	})
	app.Static("/", dir)
	app.Listen(pterm.Sprintf("127.0.0.1:%d", port))
}
