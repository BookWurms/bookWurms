package main

import (
	"log"
	"net/http"
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
	 initialize gin server context and register handlers
	*/
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB size constraint for post

	router.GET("/api/health", database.GetOk) // for testing
	router.GET("api/books", database.GetBooks(gormDB))
	router.POST("api/books", database.UploadBook(gormDB))

	if err := runServer(":50000", router); err != nil {
		log.Fatal(err)
	}
}
