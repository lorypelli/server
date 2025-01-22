package internal

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
)

func Time(dir, ext string, extension, realtime bool) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
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
			body = []byte(strings.ReplaceAll(string(body), "</body>", pterm.Sprintf("<script>new WebSocket('ws://%s:%d').onmessage=e=>e.data=='reload'&&location.reload()</script></body>", strings.Split(ctx.Hostname(), ":")[0], WS_PORT)))
			ctx.Set("Content-Type", "text/html")
			return ctx.Send(body)
		}
		return ctx.Next()
	}
}
