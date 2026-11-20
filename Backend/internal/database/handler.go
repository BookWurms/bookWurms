package database

import (
	"archive/zip"
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
