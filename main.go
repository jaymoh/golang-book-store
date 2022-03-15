package main

import (
	"github.com/gin-gonic/gin"
	"hackinroms.com/books/controllers/books"
	userscontroller "hackinroms.com/books/controllers/users"
	"hackinroms.com/books/models"
)

func main() {
	r := gin.Default()

	// connect DB instance
	models.ConnectDatabase()

	// route definitions
	r.GET("/books", bookscontroller.Index)
	r.POST("/books", bookscontroller.Store)
	r.GET("/books/:id", bookscontroller.Show)
	r.PUT("/books/:id", bookscontroller.Update)
	r.DELETE("/books/:id", bookscontroller.Destroy)

	r.GET("/users", userscontroller.Index)

	err := r.Run()
	if err != nil {
		return
	}
}
