package internal

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lorypelli/server/frontend"
)

func Path(dir string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
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
			ctx.Status(404)
			return Render(ctx, frontend.Error(ctx.Path()))
		}
		return Render(ctx, frontend.Index(ctx.Path(), p))
	}
}
