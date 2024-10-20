package main

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	"github.com/pterm/pterm"
)

func Start(dir, name string, extension bool, port int) {
	app := fiber.New(fiber.Config{
		AppName:      name,
		ServerHeader: name,
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
	app.Use(logger.New(), func(ctx *fiber.Ctx) error {
		t := ctx.Query("t")
		time := pterm.Sprint(time.Now().Unix())
		path := ctx.Path()
		file := strings.TrimPrefix(path, "/")
		if file == "" {
			file += "index.html"
		} else if !strings.Contains(file, ".") && !extension {
			file += ".html"
		}
		if t != time {
			return ctx.Redirect(pterm.Sprintf("%s?t=%s", path, time))
		}
		if strings.HasSuffix(file, ".html") {
			body, err := os.ReadFile(file)
			if err != nil {
				return ctx.Next()
			}
			body = []byte(strings.ReplaceAll(string(body), "</body>", "<script>new WebSocket('ws://127.0.0.1:50643').onmessage=e=>{e.data=='reload'&&location.reload()}</script></body>"))
			ctx.Set("Content-Type", "text/html")
			return ctx.Send(body)
		}
		return ctx.Next()
	})
	app.Static("/", dir)
	if err := app.Listen(pterm.Sprintf("127.0.0.1:%d", port)); err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
}

func StartWebsocket(dir string) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Get("/", websocket.New(func(ctx *websocket.Conn) {
		defer ctx.Close()
		hasChanged := make(chan bool)
		go Monitor(dir, hasChanged)
		if <-hasChanged {
			ctx.WriteMessage(websocket.TextMessage, []byte("reload"))
		}
	}))
	if err := app.Listen("127.0.0.1:50643"); err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
}
