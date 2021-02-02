## A Simple Go Server

This is a very simple Go server built using Fiber which performs simple CRUD operations and stores the title, author and rating of a book.

- Database used : SQLite
- Routes : 
    - Get All Books - GET /api/v1/books
    - Get Book by ID - GET /api/v1/single/:id
    - Add a book - POST /api/v1/book
        - Body - 
        ``{
            "title" : String,
            "author" : String,
            "rating" : int
        }``
    - Delete book by ID - DELETE /api/v1/delete/:id

- Steps to Run
    After cloning run -
        <br>``go install``<br>
        ``go run main.go``