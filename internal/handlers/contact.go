package handlers

import (
    "github.com/gofiber/fiber/v2"
)

type ContactHandler struct {
    BaseHandler
}

func NewContactHandler() *ContactHandler {
    return &ContactHandler{}
}

func (h *ContactHandler) SetupRoutes(app *fiber.App) {
    contact := app.Group("/api/v1/contact")
    
    contact.Get("/", h.Read())
    contact.Get("/:id", h.Read())
    contact.Post("/", h.Create())
    contact.Put("/:id", h.Update())
    contact.Delete("/:id", h.Delete())
}