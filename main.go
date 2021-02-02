package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sbiswas2209/tutorialGo/book"
	"github.com/sbiswas2209/tutorialGo/database"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect!!!")
	}
	fmt.Println("Database Connection Successful")

	database.DBConn.AutoMigrate(&book.Book{})

	fmt.Println("Database Migrated")

}

func root(c *fiber.Ctx) {
	c.Send("Server is currently up and running. Let's Go!!!")
}

func setupRoutes(a *fiber.App) {
	a.Get("/", root)
	a.Get("/api/v1/books", book.GetAllBooks)
	a.Get("/api/v1/single/:id", book.GetSingleBook)
	a.Post("/api/v1/book", book.AddBook)
	a.Delete("/api/v1/delete/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)
	app.Listen(8080)
}
