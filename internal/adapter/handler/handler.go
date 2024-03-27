package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type FiberHandler interface {
	Greet(ctx *fiber.Ctx) error
}

type fiberHandler struct {
	log *slog.Logger
}

func NewFiberHandler(i *do.Injector) (FiberHandler, error) {
	log := do.MustInvoke[*slog.Logger](i)
	fiberHandler := fiberHandler{log: log}
	return &fiberHandler, nil
}

func (h *fiberHandler) Greet(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
