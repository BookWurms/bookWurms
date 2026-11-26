package main

import (
	"log"
	"net/http"
	"src/internal/database"

	"github.com/gin-gonic/gin"

	"src/internal/handlers"
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

	if err := database.MigrateUsers(gormDB); err != nil {
		log.Fatal(err)
	}

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

	authHandler := handlers.NewAuthHandler(gormDB)

	authRoutes := router.Group("/api/auth")
	authRoutes.POST("/register", authHandler.Register)

	router.GET("/api/health", database.GetOk) // for testing
	router.GET("api/books", database.GetBooks(gormDB))
	router.POST("api/books", database.UploadBook(gormDB))

	//get all data sets by one author, to be searched for, from table books
	router.GET("/api/books/author", database.GetBooksByAuthor(gormDB))

	//get all fields from structure Book that are searchable for dropdown list
	router.GET("/api/books/searchableFields", database.GetFieldsTagedSearchable())

	//get books that match text in chosen field
	router.GET("/api/books/search", database.SearchBooklistInAnyField(gormDB))

	if err := runServer(":50000", router); err != nil {
		log.Fatal(err)
	}
}
