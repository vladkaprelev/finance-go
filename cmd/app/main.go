package main

import (
	"context"
	"log"
	"time"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/vladkaprelev/finance-go/pool"
)

func main() {
	err := godotenv.Load() // пытается загрузить .env из корня проекта
	if err != nil {
		log.Println("No .env file found or error reading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool.InitPool(ctx)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World 5!"))
	})
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Println(err.Error())
	}
}
