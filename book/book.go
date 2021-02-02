package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

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
