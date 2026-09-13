package internal

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/pterm/pterm"
)

func Logger() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		err := ctx.Next()
		status := ctx.Response().StatusCode()
		msg := fmt.Sprintf("%s (%d): %s - %s (IP: %s)", time.Now().Format(time.DateTime), status, ctx.Method(), ctx.Path(), ctx.IP())
		switch {
		case status >= 500:
			pterm.Error.Println(msg)
		case status >= 400:
			pterm.Warning.Println(msg)
		case status >= 300:
			pterm.Info.Println(msg)
		default:
			pterm.Success.Println(msg)
		}
		return err
	}
}
