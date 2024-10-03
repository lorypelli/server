package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
)

func Start(dir, name string, extension bool, port uint16) {
	app := fiber.New(fiber.Config{
		AppName:      name,
		ServerHeader: name,
	})
	if !extension {
		app.Use(func (ctx *fiber.Ctx) error {
			path := ctx.Path()
			if !strings.Contains(path, ".") {
				if err := ctx.SendFile(dir + path + ".html"); err == nil {
					return nil
				}
			}
			return ctx.Next()
		})
	}
	app.Static("/", dir)
	app.Listen(pterm.Sprintf("127.0.0.1:%d", port))
}
