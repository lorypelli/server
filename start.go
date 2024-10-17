package main

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pterm/pterm"
)

func Start(dir, name string, extension bool, port int) {
	app := fiber.New(fiber.Config{
		AppName:      name,
		ServerHeader: name,
	})
	app.Use(logger.New(), func(ctx *fiber.Ctx) error {
		path := ctx.Path()
		t := ctx.Query("t")
		time := pterm.Sprint(time.Now().Unix())
		if t != time {
			return ctx.Redirect(pterm.Sprintf("%s?t=%s", path, time))
		}
		return ctx.Next()
	})
	if !extension {
		app.Use(func(ctx *fiber.Ctx) error {
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
	if err := app.Listen(pterm.Sprintf("127.0.0.1:%d", port)); err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
}
