package routes

import (
	"github.com/gin-gonic/gin"
	bookscontroller "hackinroms.com/books/controllers/books"
	userscontroller "hackinroms.com/books/controllers/users"
)

func Setup() *gin.Engine {
	r := gin.Default()

	// route definitions
	r.GET("/books", bookscontroller.Index)
	r.POST("/books", bookscontroller.Store)
	r.GET("/books/:id", bookscontroller.Show)
	r.PUT("/books/:id", bookscontroller.Update)
	r.DELETE("/books/:id", bookscontroller.Destroy)

	r.GET("/users", userscontroller.Index)

	return r
}
