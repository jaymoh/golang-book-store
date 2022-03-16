package userscontroller

import (
	"github.com/gin-gonic/gin"
	"hackinroms.com/books/database"
	"hackinroms.com/books/models"
	"net/http"
)

func Index(ctx *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
