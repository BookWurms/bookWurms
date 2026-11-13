package database

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// simple handler that returns all books in the database in formatted JSON
func GetBooks(db *gorm.DB) gin.HandlerFunc {
	var x gin.HandlerFunc = func(context *gin.Context) {
		var b []Book
		db.Find(&b)
		context.IndentedJSON(http.StatusOK, b)
	}

	return x
}

// for testing purposes, a simple handler that returns a 200 OK response
func GetOk(context *gin.Context) {
	context.JSON(200, gin.H{"status": "ok"})
}
