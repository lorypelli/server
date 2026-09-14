package pkg

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/basicauth"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/lorypelli/server/internal"
	"github.com/pterm/pterm"
)

const shutdownTimeout = 5 * time.Second

func Start(o Options) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	app := fiber.New(fiber.Config{
		AppName:      o.Name,
		ServerHeader: o.Name,
		ErrorHandler: internal.ErrorHandler,
	})
	app.Hooks().OnPreShutdown(func() error {
		pterm.Println()
		internal.Info.Log("Shutting down...")
		return nil
	})
	app.Use(internal.Logger())
	if o.authenticated() {
		app.Use(basicAuth(o.Username, o.Password))
	}
	app.Use(internal.NoCache())
	app.Use(internal.Errors())
	app.Use(recover.New())
	app.Use(internal.TrailingSlash(o.Dir))
	app.Use(static.New(o.Dir, static.Config{
		IndexNames: []string{withExt("index", o.Ext)},
		ByteRange:  true,
	}))
	if !o.Extension {
		app.Use(implicitExtension(o.Dir, o.Ext))
	}
	app.Use(internal.Listing(o.Dir))
	announce(o)
	listen(ctx, app, o.Network, o.Port)
	internal.Success.Log("Server stopped")
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
		if path.Ext(route) == "" {
			if err := ctx.SendFile(filepath.Join(dir, withExt(route, ext))); err == nil {
				return nil
			}
		}
		return ctx.Next()
	}
}

func withExt(name, ext string) string {
	return fmt.Sprint(name, ext)
}

func announce(o Options) {
	rows := [][2]string{{"Local", internal.Link(httpURL(internal.LocalIP, o.Port))}}
	if o.Network {
		if ip := internal.GetLocalIP(); ip != internal.LocalIP {
			rows = append(rows, [2]string{"Network", internal.Link(httpURL(ip, o.Port))})
		}
	}
	rows = append(rows,
		[2]string{"", ""},
		[2]string{"Directory", absolute(o.Dir)},
		[2]string{"Extension", extensionSummary(o)},
		[2]string{"Auth", toggle(o.authenticated(), fmt.Sprintf("basic as %s", internal.Strong(o.Username)))},
	)
	lines := make([]string, len(rows))
	for i, row := range rows {
		lines[i] = fmt.Sprintf("%s  %s", internal.Muted(fmt.Sprintf("%-9s", row[0])), row[1])
	}
	pterm.DefaultBox.WithTitle(internal.Strong(o.title())).WithTitleTopCenter().Println(strings.Join(lines, "\n"))
	pterm.Println(internal.Muted("Press Ctrl+C to stop"))
	pterm.Println()
}

func absolute(dir string) string {
	if abs, err := filepath.Abs(dir); err == nil {
		return abs
	}
	return dir
}

func extensionSummary(o Options) string {
	if o.Extension {
		return fmt.Sprintf("%s %s", o.Ext, internal.Muted("(explicit in URLs)"))
	}
	return fmt.Sprintf("%s %s", o.Ext, internal.Muted("(implicit, resolved automatically)"))
}

func toggle(enabled bool, detail string) string {
	if enabled {
		return pterm.Green(detail)
	}
	return internal.Muted("disabled")
}

func httpURL(host string, port uint16) string {
	return (&url.URL{Scheme: "http", Host: address(host, port)}).String()
}

func address(host string, port uint16) string {
	return net.JoinHostPort(host, strconv.Itoa(int(port)))
}

func listen(ctx context.Context, app *fiber.App, network bool, port uint16) {
	host := internal.LocalIP
	if network {
		host = ""
	}
	addr := address(host, port)
	err := app.Listen(addr, fiber.ListenConfig{
		DisableStartupMessage: true,
		GracefulContext:       ctx,
		ShutdownTimeout:       shutdownTimeout,
	})
	if err != nil {
		internal.Exit(fmt.Errorf("cannot listen on %s: %w", addr, err))
	}
}
