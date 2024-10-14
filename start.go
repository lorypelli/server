package main

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
)

func Start(dir, name string, extension bool, port uint16) {
	app := fiber.New(fiber.Config{
		AppName:      name,
		ServerHeader: name,
	})
	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		path := ctx.Path()
		t := ctx.Query("t")
		time := time.Now().Unix()
		if !extension {
			if !strings.Contains(path, ".") {
				if err := ctx.SendFile(dir + path + ".html"); err == nil {
					return nil
				}
			}
		}
		if strings.TrimSpace(t) == "" || t < pterm.Sprint(time) {
			return ctx.Redirect(pterm.Sprintf("%s?t=%d", path, time))
		}
		return ctx.Next()
	})
	app.Static("/", dir)
	if err := app.Listen(pterm.Sprintf("127.0.0.1:%d", port)); err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
}
