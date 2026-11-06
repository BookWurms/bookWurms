package main

import (
	"fmt"
	"time"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
)

func main() {

	// Simple "Hello, World!" to test the docker container setup
	for true {
		fmt.Println("Hello, World!")
		time.Sleep(1 * time.Second)
	}
	_ = chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":3000", r)
}
