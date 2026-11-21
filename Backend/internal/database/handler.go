package database

import (
	"archive/zip"
	"log"

	"fmt"
	"net/http"
	"path/filepath"
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
		fmt.Println(authorToBeSearchedFor)

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

		upload, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File could not be read"})
			return
		}
		defer upload.Close()

		//check if its actually a compressed zip file
		// without this, a file renamed as .epub would cause gin to panic.
		readZip, err := zip.NewReader(upload, file.Size)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ".epub file is invalid"})
			return
		}

		//ok so now we can take the first file, which in an epub should always be mimetype, and the
		//exact byte size should be 20. if this isnt true, we send a bad request. Only then, do
		// we actually attempt saving and seeding.
		// https://www.edrlab.org/open-standards/anatomy-of-an-epub-3-file/
		// This is still vulnerable, but for our use-case, we can either leave this as is, or go further.
		firstLine := readZip.File[0]
		if firstLine.Name == "mimetype" && firstLine.UncompressedSize64 == 20 {
			fmt.Printf("Filename: %s || Size: %d \n", readZip.File[0].Name, firstLine.UncompressedSize64)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": ".epub file is invalid"})
			return
		}

		c.SaveUploadedFile(file, dst) //physically save this to the /books after all parsing is done.
		SeedBook(db, dst)
		//debug print
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded successfully", file.Filename))
	}
}

// get all books that match text within field
func SearchBooklistInAnyField(datasbase *gorm.DB) gin.HandlerFunc {
	return func(webpage_context *gin.Context) {

		var booklist []Book

		field := webpage_context.Query("field")
		text := webpage_context.Query(("query_author"))
		log.Println("140", field, text)
		// datasbase.Where(field+" = ?", "Fyodor Dostoevsky").Find(&booklist)
		datasbase.Where(field+"= ?", text).Find(&booklist)
		//format result to JSON)

		webpage_context.JSON(200, booklist)
	}
}
