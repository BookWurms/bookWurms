package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"src/internal/database"

	"github.com/gin-gonic/gin"
)

func runServer(addr string, handler http.Handler) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return srv.ListenAndServe()
}

func main() {
	gormDB, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	if err := database.CreateSeededBookTable(gormDB); err != nil {
		log.Fatal(err)
	}

	if err := database.SeedAllBooks(gormDB, "/books"); err != nil {
		log.Fatal(err)
	}

	/*
	 initialize gin server context
	*/
	router := gin.Default()

	router.GET("/api/health", database.GetOk) // for testing

	router.GET("api/books", database.GetBooks(gormDB))
	// router.POST("/api/books", database.UploadBook(gormDB))

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("api/books", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("ebook")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(file.Filename)

		// Upload the file to specific dst.
		dst := filepath.Join("/books", filepath.Base(file.Filename))
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		database.SeedBook(gormDB, dst)
	})

	if err := runServer(":50000", router); err != nil {
		log.Fatal(err)
	}
}
