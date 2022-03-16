package bookscontroller

import (
	"github.com/gin-gonic/gin"
	"hackinroms.com/books/database"
	"hackinroms.com/books/models"
	"net/http"
)

// Index get all books
func Index(ctx *gin.Context) {
	var books []models.Book

	database.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func Show(ctx *gin.Context) {
	var book models.Book

	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func Store(ctx *gin.Context) {
	var input BookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store book
	book := models.Book{Title: input.Title, Author: input.Author}
	database.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func Update(ctx *gin.Context) {
	// retrieve model if it exists
	var book models.Book
	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	// validate input
	var input UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&book).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func Destroy(ctx *gin.Context) {
	// retrieve model if it exists
	var book models.Book
	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	database.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": "Book Deleted"})
}
