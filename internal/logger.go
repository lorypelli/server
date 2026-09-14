package internal

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/lorypelli/server/frontend/utils"
	"github.com/pterm/pterm"
)

func Logger() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		start := time.Now()
		err := ctx.Next()
		status := ctx.Response().StatusCode()
		Badge(fmt.Sprint(status), statusStyle(status)).Log("%s", Columns(
			Muted(fmt.Sprintf("%8s", Duration(time.Since(start)))),
			Muted(fmt.Sprintf("%9s", size(ctx))),
			Muted(fmt.Sprintf("%-15s", ctx.IP())),
			fmt.Sprintf("%s %s", Strong(fmt.Sprintf("%-7s", ctx.Method())), displayPath(ctx)),
		))
		return err
	}
}

func statusStyle(status int) *pterm.Style {
	switch {
	case status >= fiber.StatusInternalServerError:
		return &pterm.ThemeDefault.ErrorPrefixStyle
	case status >= fiber.StatusBadRequest:
		return &pterm.ThemeDefault.WarningPrefixStyle
	case status >= fiber.StatusMultipleChoices:
		return &pterm.ThemeDefault.InfoPrefixStyle
	default:
		return &pterm.ThemeDefault.SuccessPrefixStyle
	}
}

func size(ctx fiber.Ctx) string {
	res := ctx.Response()
	if !res.IsBodyStream() {
		return utils.Size(int64(len(res.Body())))
	}
	if n := res.Header.ContentLength(); n >= 0 {
		return utils.Size(int64(n))
	}
	return "-"
}
