package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"src/internal/database"
	"src/internal/handlers"
)

const miB = 1 << 20
const maxMemory = 8 * miB

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
	router.MaxMultipartMemory = maxMemory

	authHandler := handlers.NewAuthHandler(gormDB)

	/* API routes */
	apiRoutes := router.Group("/api")
	apiRoutes.GET("/health", database.GetOk) // docker health check

	/* Authentication routes */
	authRoutes := apiRoutes.Group("/auth")
	authRoutes.POST("/register", authHandler.Register)

	/* Book routes */
	bookRoutes := apiRoutes.Group("/books")
	bookRoutes.GET("", database.GetBooks(gormDB))
	bookRoutes.POST("", database.UploadBook(gormDB))
	bookRoutes.GET("/author", database.GetBooksByAuthor(gormDB))             // by one author
	bookRoutes.GET("/searchableFields", database.GetFieldsTagedSearchable()) //dropdown list, searchable books
	bookRoutes.GET("/search", database.SearchBooklistInAnyField(gormDB))     //match chosen text field

	if err := runServer(":50000", router); err != nil {
		log.Fatal(err)
	}
}
