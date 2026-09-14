package internal

import (
	"errors"
	"io/fs"
	"net/url"

	"github.com/gofiber/fiber/v3"
	"github.com/lorypelli/server/frontend"
	"github.com/lorypelli/server/frontend/utils"
)

func Errors() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		return ErrorHandler(ctx, ctx.Next())
	}
}

func ErrorHandler(ctx fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	status := fiber.StatusInternalServerError
	if e, ok := errors.AsType[*fiber.Error](err); ok {
		status = e.Code
	} else {
		Error.Log("%s %s: %v", ctx.Method(), displayPath(ctx), err)
	}
	ctx.Status(status)
	return render(ctx, frontend.Error(displayPath(ctx), status, utils.ParseView(ctx.Query("view"))))
}

func displayPath(ctx fiber.Ctx) string {
	route := ctx.Path()
	if unescaped, err := url.PathUnescape(route); err == nil {
		return unescaped
	}
	return route
}

func classify(err error) error {
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return fiber.ErrNotFound
	case errors.Is(err, fs.ErrPermission):
		return fiber.ErrForbidden
	default:
		return err
	}
}
