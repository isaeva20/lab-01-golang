package handlers

import (
    "github.com/gofiber/fiber/v2"
)

type CRUDHandler interface {
    Create() fiber.Handler
    Read() fiber.Handler
    Update() fiber.Handler
    Delete() fiber.Handler
}

type BaseHandler struct{}

func (h *BaseHandler) Create() fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        return ctx.SendString("Create")
    }
}

func (h *BaseHandler) Read() fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        return ctx.SendString("Read")
    }
}

func (h *BaseHandler) Update() fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        return ctx.SendString("Update")
    }
}

func (h *BaseHandler) Delete() fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        return ctx.SendString("Delete")
    }
}