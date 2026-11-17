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

// nsloniow
// return all books into JSON from table books for one particulare author
func GetBooksByAuthor(myDatabase *gorm.DB) gin.HandlerFunc {

	// //printing 'cool'
	// return func(context *gin.Context) {
	// 	context.JSON(http.StatusOK, "cool")
	// }

	//select from db
	return func(context *gin.Context) {
		var booksByAuthor []Book
		//SELECT * FROM books WHERE author = "Fyodor";
		myDatabase.Where("author = ?", "Fyodor Dostoevsky").Find(&booksByAuthor)
		context.IndentedJSON(http.StatusOK, booksByAuthor)
	}
}
