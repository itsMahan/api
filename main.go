package main

import (
	"errors"
	"flag"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//////////---Types---//////////

type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Count  int    `json:"count"`
}

var books = []book{
	{ID: 1, Title: "Things Fall Apart", Author: "Chinua Achebe", Count: 12},
	{ID: 2, Title: "Pride and Prejudice", Author: "Jane Austen", Count: 7},
	{ID: 3, Title: "The Kite Runner", Author: "Khaled Hosseini", Count: 21},
	{ID: 4, Title: "War and Peace", Author: "Leo Tolstoy", Count: 3},
	{ID: 5, Title: "Silent Spring", Author: "Rachel Carson", Count: 14},
	{ID: 6, Title: "1984", Author: "George Orwell", Count: 5},
	{ID: 7, Title: "Gone Girl", Author: "Gillian Flynn", Count: 11},
}

//////////---Handlers---//////////

func GetAllBooksHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBookHandler(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook.ID = len(books) + 1

	books = append(books, newBook)

	c.JSON(http.StatusCreated, newBook)
}

func getBookByIdHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := findBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Message: ": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, *book)
}

func getBookByTitleHandler(c *gin.Context) {
	title := c.Param("title")
	book, err := findBookByTitle(title)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Message: ": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, *book)
}

func checkOutBookByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	amountStr := c.Query("amount")

	id, err_id := strconv.Atoi(idStr)
	amount, err_amount := strconv.Atoi(amountStr)

	if err_id != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is invalid"})
		return
	}

	if err_amount != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount parameter is invalid"})
		return
	}

	book, err := findBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Message: ": err.Error()})
		return
	}

	if book.Count-amount < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Message: ": "this amount of the book is not available."})
		return
	}

	book.Count -= amount

	c.IndentedJSON(http.StatusOK, book)
}

func CheckInBookByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	amountStr := c.Query("amount")

	id, err_id := strconv.Atoi(idStr)
	amount, err_amount := strconv.Atoi(amountStr)

	if err_id != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is invalid"})
		return
	}

	if err_amount != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount parameter is invalid"})
		return
	}

	book, err := findBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Message: ": err.Error()})
		return
	}

	book.Count += amount

	c.IndentedJSON(http.StatusOK, book)
}

//////////---Functions---//////////

func findBookById(id int) (*book, error) {
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func findBookByTitle(name string) (*book, error) {
	for i, book := range books {
		if book.Title == name {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func main() {
	Addr := flag.String("addr", ":8080", "Listen and Serve Address")
	flag.Parse()

	engine := gin.Default()

	engine.GET("/allbooks", GetAllBooksHandler)
	engine.GET("/book/id/:id", getBookByIdHandler)
	engine.GET("/book/title/:title", getBookByTitleHandler)
	engine.POST("/addBook", addBookHandler)
	engine.PATCH("/book/check_out/:id", checkOutBookByIdHandler)
	engine.PATCH("/book/check_in/:id", CheckInBookByIdHandler)

	engine.Run(*Addr)
}
