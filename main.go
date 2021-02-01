package main

import (
	"github.com/gofiber/fiber"
	"github.com/sbiswas2209/tutorialGo/book"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World")
}

func setupRoutes(a *fiber.App) {
	a.Get("/api/v1/books", book.GetAllBooks)
	a.Get("/api/v1/book/:id", book.GetSingleBook)
	a.Post("/api/v1/book", book.AddBook)
	a.Delete("/api/v1/book/delete/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(8080)
}
