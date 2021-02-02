package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/sbiswas2209/tutorialGo/database"
)

//Book is struct for all data stored in books.db
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

//GetAllBooks is for Get all books
func GetAllBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

//GetSingleBook is for Get books by id
func GetSingleBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		c.Status(404).Send("No Book found at given Index")
		return
	}
	c.JSON(book)
}

//AddBook is for Add book
func AddBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)

	c.JSON(book)
}

//DeleteBook is for Delete book
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.Status(404).Send("No Book found with given ID")
		return
	}
	db.Delete(&book)
	c.Status(200).Send("Book Successfully deleted")
}
