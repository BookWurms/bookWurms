package database

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

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

/*
Migrated from main, this will take the file, and join it into the uploads folder.
Needs to validate the file is in fact an epub in the backend.
*/
func UploadBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("ebook")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(file.Filename)

		if filepath.Ext(file.Filename) != ".epub" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file is not .epub"})
			return
		}

		dst := filepath.Join("/books", filepath.Base(file.Filename))
		SeedBook(db, dst)

		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}
