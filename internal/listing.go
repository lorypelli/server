package internal

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/lorypelli/server/frontend"
	"github.com/lorypelli/server/frontend/utils"
)

func Listing(dir string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		route := ctx.Path()
		unescaped, err := url.PathUnescape(route)
		if err != nil {
			return fiber.ErrBadRequest
		}
		target := filepath.Join(dir, unescaped)
		info, err := os.Stat(target)
		if err != nil {
			return Classify(err)
		}
		if !info.IsDir() {
			return fiber.ErrForbidden
		}
		entries, err := utils.Entries(target)
		if err != nil {
			return Classify(err)
		}
		return Render(ctx, frontend.Index(route, entries))
	}
}
