package pkg

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	"github.com/lorypelli/server/internal"
	t "github.com/lorypelli/server/templ"
	"github.com/pterm/pterm"
)

var IP = internal.GetLocalIP()

var LocalIP = "127.0.0.1"

func Start(dir, ext, name string, extension, network, realtime bool, port, ws_port uint16) {
	app := fiber.New(fiber.Config{
		AppName:               name,
		ServerHeader:          name,
		DisableStartupMessage: true,
	})
	if !network {
		IP = LocalIP
	}
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
	app.Use(logger.New(), func(ctx *fiber.Ctx) error {
		t := ctx.Query("t")
		time := pterm.Sprint(time.Now().Unix())
		path := ctx.Path()
		file := strings.TrimPrefix(path, "/")
		if file == "" {
			file += "index" + ext
		} else if !strings.Contains(file, ".") && !extension {
			file += ext
		}
		if t != time {
			return ctx.Redirect(pterm.Sprintf("%s?t=%s", path, time))
		}
		if realtime && strings.HasSuffix(file, ext) {
			body, err := os.ReadFile(file)
			if err != nil {
				return ctx.Next()
			}
			body = []byte(strings.ReplaceAll(string(body), "</body>", pterm.Sprintf("<script>new WebSocket('ws://%s:%d').onmessage=e=>e.data=='reload'&&location.reload()</script></body>", IP, ws_port)))
			ctx.Set("Content-Type", "text/html")
			return ctx.Send(body)
		}
		return ctx.Next()
	})
	app.Static("/", dir, fiber.Static{
		Index: "index" + ext,
	})
	app.Use(func(ctx *fiber.Ctx) error {
		path := strings.Split(ctx.Path(), "/")
		var p string
		if len(path) < 2 {
			p = dir
		} else {
			p = dir + strings.Join(path[1:], "/")
		}
		abs, err := filepath.Abs(p)
		if err != nil {
			return ctx.Next()
		}
		return internal.Render(ctx, t.Index(abs))
	})
	box := pterm.DefaultBox.WithTitle(name).WithTitleTopCenter()
	if IP != LocalIP {
		box.Printfln("Local: http://%s:%d\nNetwork: http://%s:%d", LocalIP, port, IP, port)
		if err := app.Listen(pterm.Sprintf(":%d", port)); err != nil {
			internal.Exit(err)
		}
	} else {
		box.Printfln("Local: http://%s:%d", LocalIP, port)
		if err := app.Listen(pterm.Sprintf("%s:%d", LocalIP, port)); err != nil {
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
	if IP != LocalIP {
		if err := app.Listen(pterm.Sprintf(":%d", port)); err != nil {
			internal.Exit(err)
		}
	} else {
		if err := app.Listen(pterm.Sprintf("%s:%d", LocalIP, port)); err != nil {
			internal.Exit(err)
		}
	}
}
