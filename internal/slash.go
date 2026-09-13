package internal

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func TrailingSlash(dir string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		route := ctx.Path()
		if strings.HasSuffix(route, "/") {
			return ctx.Next()
		}
		unescaped, err := url.PathUnescape(route)
		if err != nil {
			return ctx.Next()
		}
		if info, err := os.Stat(filepath.Join(dir, unescaped)); err != nil || !info.IsDir() {
			return ctx.Next()
		}
		target := fmt.Sprintf("%s/", route)
		return ctx.Redirect().Status(fiber.StatusFound).To(target)
	}
}
