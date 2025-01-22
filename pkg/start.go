package pkg

import (
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/websocket/v2"
	"github.com/lorypelli/server/internal"
	"github.com/pterm/pterm"
)

var IP = internal.GetLocalIP()

func Start(dir, ext, name, username, password string, extension, network, realtime bool, port uint16) {
	app := fiber.New(fiber.Config{
		AppName:               name,
		ServerHeader:          name,
		DisableStartupMessage: true,
	})
	if !network {
		IP = internal.LOCAL_IP
	}
	if username != "" && password != "" {
		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				username: password,
			},
		}))
	}
	app.Use(internal.Logger())
	app.Use(internal.Time(dir, ext, extension, realtime))
	app.Static("/", dir, fiber.Static{
		Index: "index" + ext,
		ModifyResponse: func(ctx *fiber.Ctx) error {
			ctx.Set("Content-Disposition", pterm.Sprintf("inline; filename=%q", filepath.Base(ctx.Path())))
			ctx.Set("Content-Length", pterm.Sprint(len(ctx.Response().Body())))
			return nil
		},
	})
	app.Use(internal.Path(dir))
	if !extension {
		app.Use(func(ctx *fiber.Ctx) error {
			path := ctx.Path()
			if !strings.Contains(path, ".") {
				if err := ctx.SendFile(dir + path + ext); err == nil {
					return nil
				}
			}
			return ctx.Next()
		})
	}
	box := pterm.DefaultBox.WithTitle(name).WithTitleTopCenter()
	msg := pterm.Sprintf("Local: http://%s:%d", internal.LOCAL_IP, port)
	if IP != internal.LOCAL_IP {
		msg += "\n"
		msg += pterm.Sprintf("Network: http://%s:%d", IP, port)
	}
	box.Println(msg)
	if IP != internal.LOCAL_IP {
		if err := app.Listen(pterm.Sprintf(":%d", port)); err != nil {
			internal.Exit(err)
		}
	} else {
		if err := app.Listen(pterm.Sprintf("%s:%d", internal.LOCAL_IP, port)); err != nil {
			internal.Exit(err)
		}
	}
}

func StartWebsocket(dir string, port uint16) {
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
	if IP != internal.LOCAL_IP {
		if err := app.Listen(pterm.Sprintf(":%d", port)); err != nil {
			internal.Exit(err)
		}
	} else {
		if err := app.Listen(pterm.Sprintf("%s:%d", internal.LOCAL_IP, port)); err != nil {
			internal.Exit(err)
		}
	}
}
