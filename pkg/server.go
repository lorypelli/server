package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/basicauth"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/lorypelli/server/internal"
	"github.com/pterm/pterm"
)

func Start(o Options) {
	if o.Realtime {
		go serveWebsocket(o.Dir, o.Network)
	}
	app := fiber.New(fiber.Config{
		AppName:      o.Name,
		ServerHeader: o.Name,
		ErrorHandler: internal.ErrorHandler,
	})
	if o.Username != "" && o.Password != "" {
		app.Use(basicAuth(o.Username, o.Password))
	}
	app.Use(internal.Logger())
	app.Use(internal.NoCache())
	if o.Realtime {
		app.Use(internal.Realtime())
	}
	app.Use(internal.Errors())
	app.Use(recover.New())
	app.Use(internal.TrailingSlash(o.Dir))
	app.Use(static.New(o.Dir, static.Config{
		IndexNames: []string{fmt.Sprintf("index%s", o.Ext)},
		ByteRange:  true,
	}))
	if !o.Extension {
		app.Use(implicitExtension(o.Dir, o.Ext))
	}
	app.Use(internal.Listing(o.Dir))
	announce(o.Name, o.Network, o.Port)
	listen(app, o.Network, o.Port)
}

func basicAuth(username, password string) fiber.Handler {
	sum := sha256.Sum256([]byte(password))
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			username: fmt.Sprintf("{SHA256}%s", base64.StdEncoding.EncodeToString(sum[:])),
		},
	})
}

func implicitExtension(dir, ext string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		route := ctx.Path()
		if !strings.Contains(route, ".") {
			if err := ctx.SendFile(filepath.Join(dir, fmt.Sprintf("%s%s", route, ext))); err == nil {
				return nil
			}
		}
		return ctx.Next()
	}
}

func announce(name string, network bool, port uint16) {
	msg := fmt.Sprintf("Local: http://%s:%d", internal.LocalIP, port)
	if network {
		if ip := internal.GetLocalIP(); ip != internal.LocalIP {
			msg = fmt.Sprintf("%s\nNetwork: http://%s:%d", msg, ip, port)
		}
	}
	pterm.DefaultBox.WithTitle(name).WithTitleTopCenter().Println(msg)
}

func listen(app *fiber.App, network bool, port uint16) {
	host := internal.LocalIP
	if network {
		host = ""
	}
	if err := app.Listen(fmt.Sprintf("%s:%d", host, port), fiber.ListenConfig{DisableStartupMessage: true}); err != nil {
		internal.Exit(err)
	}
}
