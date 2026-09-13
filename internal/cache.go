package internal

import "github.com/gofiber/fiber/v3"

func NoCache() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		err := ctx.Next()
		ctx.Set(fiber.HeaderCacheControl, "no-store")
		return err
	}
}
