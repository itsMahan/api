# Simple Book Management API

This project is a Simple RESTful API for managing a collection of books. It allows users to perform CRUD operations on the books, such as adding, retrieving, checking out, and checking in books. The API is built using the Gin framework in Go.

## Features

- Retrieve all books in the collection.
- Retrieve a book by its ID or title.
- Add a new book to the collection.
- Check out a specified amount of a book by its ID.
- Check in  a specified amount of a book by its ID.

## Installation

### Prerequisites

- Go (https://golang.org/dl/)
- Gin (https://github.com/gin-gonic/gin)

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/book-management-api.git
   cd book-management-api

2. Install the dependencies:
   go get -u github.com/gin-gonic/gin

3. Build and run the application:
   go build ./main.go
   ./main.go -addr=:8080

### Usage

The API can be accessed at http://localhost:8080

---ENDPOINTS---

-Get All Books:
	GET /allbooks

-Get a Book by ID:
	GET /book/id/:id

-Get a Book by Title:
	GET /book/title/:title

-Add a New Book:
	POST /addBook


-Check Out a Book by ID:
	PATCH /book/check_out/:id

-Check In a Book by ID:
	PATCH /book/check_in/:id


---Example Requests---

-Get All Books:
	curl -X GET http://localhost:8080/allbooks

-Get a Book by ID
	curl -X GET http://localhost:8080/book/id/1

-Get a Book by Title
	curl -X GET http://localhost:8080/book/title/1984

-Add a New Book
	curl -X POST http://localhost:8080/addBook -d '{"title":"48 Laws Of Power", "author":"Robert Green", "count":10}'

-Check Out a Book
	curl -X PATCH http://localhost:8080/book/check_out/1?amount=2

-Check In a Book
	curl -X PATCH http://localhost:8080/book/check_in/1?amount=2


### Note

this project is foe personal purposes only and i did it with help of some recourses in order to learn 
i know there is a lot of space to improve and develop the code
for example this projects needs a database server that i didn't know how to implement YET 
i try my best to work on it in the future 
