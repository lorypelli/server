package internal

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Realtime() fiber.Handler {
	script := fmt.Sprintf("<script>new WebSocket(`ws://${location.hostname}:%d`).onmessage=e=>e.data=='reload'&&location.reload()</script>", WSPort)
	return func(ctx fiber.Ctx) error {
		err := ctx.Next()
		if strings.HasPrefix(ctx.GetRespHeader(fiber.HeaderContentType), fiber.MIMETextHTML) && len(ctx.Response().Body()) > 0 {
			ctx.Response().AppendBodyString(script)
		}
		return err
	}
}
