package book

import (
	"github.com/gofiber/fiber"
)

func GetAllBooks(c *fiber.Ctx) {
	c.Send("Get All Books")
}

func GetSingleBook(c *fiber.Ctx) {
	c.Send("Get Single Book")
}

func AddBook(c *fiber.Ctx) {
	c.Send("Add Book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete Book")
}
