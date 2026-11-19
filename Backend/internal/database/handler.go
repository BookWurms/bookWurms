package database

import (
	"net/http"
	"reflect"

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

// return all books into JSON from table books for one particulare author
func GetBooksByAuthor(myDatabase *gorm.DB) gin.HandlerFunc {

	//select from table books, as it is structure Book,
	return func(webpage_content *gin.Context) {

		var booksByAuthor []Book

		////localhost:8443/api/books/author?query_author=Fyodor%20Dostoevsky
		authorToBeSearchedFor := webpage_content.Query("query_author")

		//print inside docker container 'backend' = where this runs
		//docker compose logs -f backend
		println("%s", authorToBeSearchedFor)

		//SELECT * FROM books WHERE author = authorToBeSearchedFor and save in booksByAuthor
		//GORM names a table created from a struct with small first letter and pluralising it
		//struct Book will be table books.
		//GORM looks at type of booksByAuthor and Books is tied to table books
		myDatabase.Where("author = ?", authorToBeSearchedFor).Find(&booksByAuthor)

		//format result, the values in booksByAuthor, into JSON
		webpage_content.IndentedJSON(http.StatusOK, booksByAuthor)
	}
}

// get all fields from structure Book that are searchable for dropdown list
func GetFieldsTagedSearchable() gin.HandlerFunc {

	return func(webpage_context *gin.Context) {
		// name of first field
		// webpage_context.JSON(200, reflect.TypeOf(Book{}).Field(1).Name)

		//all fields
		searchableFields := []string{}
		for field_count := 0; field_count < reflect.TypeOf(Book{}).NumField(); field_count++ {
			if (reflect.TypeOf(Book{}).Field(field_count).Tag.Get("searchable") == "true") {
				searchableFields = append(searchableFields, reflect.TypeOf(Book{}).Field(field_count).Name)
			}
		}
		webpage_context.JSON(200, searchableFields)
	}
}
