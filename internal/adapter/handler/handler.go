package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type FiberHandler interface {
	Greet(ctx *fiber.Ctx) error
}

type fiberHandler struct {
	log *slog.Logger
}

func NewFiberHandler(log *slog.Logger) FiberHandler {
	return &fiberHandler{log: log}
}

func (h *fiberHandler) Greet(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
