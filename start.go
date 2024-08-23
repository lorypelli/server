package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Start(dir string, port uint16) {
	app := fiber.New(fiber.Config{
		Prefork: true,
		ServerHeader: "Fiber",
	})
	app.Static("/", dir)
	app.Listen(fmt.Sprintf("127.0.0.1:%d", port))
}
