package pkg

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/lorypelli/server/frontend"
	"github.com/lorypelli/server/internal"
	"github.com/pterm/pterm"
)

var IP = internal.GetLocalIP()

func Start(dir, ext, name string, extension, network, realtime bool, port, ws_port uint16) {
	app := fiber.New(fiber.Config{
		AppName:               name,
		ServerHeader:          name,
		DisableStartupMessage: true,
	})
	if !network {
		IP = internal.LOCAL_IP
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
	app.Use(internal.Logger())
	app.Use(func(ctx *fiber.Ctx) error {
		t := pterm.Sprint(time.Now().Unix())
		path := ctx.Path()
		file := strings.TrimPrefix(path, "/")
		if file == "" {
			file += "index" + ext
		} else if !extension && !strings.Contains(file, ".") {
			file += ext
		}
		if ctx.Query("t") != t {
			return ctx.Redirect(pterm.Sprintf("%s?t=%s", path, t))
		}
		if realtime && strings.HasSuffix(file, ext) {
			body, err := os.ReadFile(dir + "/" + file)
			if err != nil {
				return ctx.Next()
			}
			body = []byte(strings.ReplaceAll(string(body), "</body>", pterm.Sprintf("<script>new WebSocket('ws://%s:%d').onmessage=e=>e.data=='reload'&&location.reload()</script></body>", strings.Split(ctx.Hostname(), ":")[0], ws_port)))
			ctx.Set("Content-Type", "text/html")
			return ctx.Send(body)
		}
		return ctx.Next()
	})
	app.Static("/", dir, fiber.Static{
		Index: "index" + ext,
		ModifyResponse: func(ctx *fiber.Ctx) error {
			ctx.Response().Header.Set("Content-Disposition", pterm.Sprintf("inline; filename=%q", filepath.Base(ctx.Path())))
			ctx.Response().Header.Set("Content-Length", pterm.Sprint(len(ctx.Response().Body())))
			return nil
		},
	})
	app.Use(func(ctx *fiber.Ctx) error {
		path := strings.Split(ctx.Path(), "/")
		p, err := filepath.Abs(dir)
		if err != nil {
			return ctx.Next()
		}
		if len(path) >= 2 {
			p += "/" + strings.Join(path[1:], "/")
		}
		p, err = filepath.Abs(p)
		if err != nil {
			return ctx.Next()
		}
		p, err = url.QueryUnescape(p)
		if err != nil {
			return ctx.Next()
		}
		if _, err := os.Stat(p); err != nil {
			return internal.Render(ctx, frontend.Error(ctx.Path()))
		}
		return internal.Render(ctx, frontend.Index(ctx.Path(), p))
	})
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
