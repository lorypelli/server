package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
)

func Logger() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		now := time.Now().Format("15:04:05")
		status := ctx.Response().StatusCode()
		method := ctx.Method()
		path := ctx.Path()
		ip := ctx.IP()
		msg := pterm.Sprintf("%s (%d): %s - %s (IP: %s)", now, status, method, path, ip)
		if status >= 200 && status < 300 {
			pterm.Success.Println(msg)
		} else if status >= 300 && status < 400 {
			pterm.Info.Println(msg)
		} else if status >= 400 && status < 500 {
			pterm.Warning.Println(msg)
		} else if status >= 500 {
			pterm.Error.Println(msg)
		}
		return err
	}
}
