package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"src/internal/database"
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

	/* API routes */
	apiRoutes := router.Group("/api")

	/* Authentication routes */
	authRoutes := apiRoutes.Group("/auth")
	authRoutes.POST("/register", authHandler.Register)

	/* General API routes */
	apiRoutes.GET("/health", database.GetOk) // for testing

	/* Book routes */
	bookRoutes := apiRoutes.Group("/books")

	bookRoutes.GET("", database.GetBooks(gormDB))
	bookRoutes.POST("", database.UploadBook(gormDB))

	// get all data sets by one author, to be searched for, from table books
	bookRoutes.GET("/author", database.GetBooksByAuthor(gormDB))

	//get all fields from structure Book that are searchable for dropdown list
	bookRoutes.GET("/searchableFields", database.GetFieldsTagedSearchable())

	//get books that match text in chosen field
	bookRoutes.GET("/search", database.SearchBooklistInAnyField(gormDB))

	if err := runServer(":50000", router); err != nil {
		log.Fatal(err)
	}
}
