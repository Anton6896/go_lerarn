package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// object struct
type Book struct {
	// allow json serialization
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Amount int    `json:"amount"`
}

// local db
var books = []Book{
	{Id: "1", Title: "title1", Author: "author1", Amount: 3},
	{Id: "2", Title: "title2", Author: "author2", Amount: 5},
	{Id: "3", Title: "title3", Author: "author3", Amount: 7},
}

// crud
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"err": err})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBookByIdUtil(id string) (*Book, error) { // helper
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("cant find this book")
}

func RetriveBook(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookByIdUtil(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"err": "cant find this book"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func CheckOutBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok { // ok is bool param
		c.IndentedJSON(http.StatusNotFound, gin.H{"err": "cant handle this id"})
		return
	}

	book, err := GetBookByIdUtil(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"err": "cant find book with this id"})
		return
	} 

	if book.Amount <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"err": "not enough book left"})
		return
	}

	book.Amount --
	c.IndentedJSON(http.StatusOK, book)
}

// main
func main() {
	log.Println("init golang local api")
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/books", GetBooks)
	router.GET("/books/:id", RetriveBook)
	router.POST("/books", CreateBook)
	router.POST("/books/checkout", CheckOutBook)

	router.Run("localhost:8080") // same as with log.Fatal
}
