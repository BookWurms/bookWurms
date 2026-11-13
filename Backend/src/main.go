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
	 initialize gin server context
	*/
	r := gin.Default()

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	if err := runServer(":50000", r); err != nil {
		log.Fatal(err)
	}
}
