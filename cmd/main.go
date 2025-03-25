package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com.isaeva20/lab-01/internal/handlers"
)

func main() {
    app := fiber.New()

    contactHandler := handlers.NewContactHandler()
    groupHandler := handlers.NewGroupHandler()

    contactHandler.SetupRoutes(app)
    groupHandler.SetupRoutes(app)

    log.Println("Сервер запущен на порту 6080")
    if err := app.Listen(":6080"); err != nil {
        log.Fatal(err)
    }
}