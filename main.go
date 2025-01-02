package main

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/vladkaprelev/finance-go/db"
)

func main() {
	err := godotenv.Load() // пытается загрузить .env из корня проекта
	if err != nil {
		log.Println("No .env file found or error reading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.NewPool(ctx)
	if err != nil {
		log.Fatalf("Ошибка при инициализации базы: %v", err)
	}
	defer pool.Close()

	log.Println("Подключение к базе PostgreSQL установлено!")
}
