package handlers

import (
    "github.com/gofiber/fiber/v2"
)

type GroupHandler struct {
    BaseHandler
}

func NewGroupHandler() *GroupHandler {
    return &GroupHandler{}
}

func (h *GroupHandler) SetupRoutes(app *fiber.App) {
    group := app.Group("/api/v1/group")
    
    group.Get("/", h.Read())
    group.Get("/:id", h.Read())
    group.Post("/", h.Create())
    group.Put("/:id", h.Update())
    group.Delete("/:id", h.Delete())
}