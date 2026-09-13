package internal

import (
	"errors"
	"io/fs"

	"github.com/gofiber/fiber/v3"
	"github.com/lorypelli/server/frontend"
	"github.com/lorypelli/server/frontend/utils"
	"github.com/pterm/pterm"
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
		pterm.Error.Println(err)
	}
	ctx.Status(status)
	return render(ctx, frontend.Error(ctx.Path(), status, utils.ParseView(ctx.Query("view"))))
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
