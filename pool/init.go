package pool

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newPool(ctx context.Context) (*pgxpool.Pool, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT") // убедитесь, что имя переменной правильно "DB_PORT", а не "DB_POTR"
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Используем net.JoinHostPort для формирования строки "host:port"
	hostPort := net.JoinHostPort(dbHost, dbPort)

	// Формируем DSN
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s", dbUser, dbPassword, hostPort, dbName)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("не удалось распарсить конфиг: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать пул соединений: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()

		return nil, fmt.Errorf("не удалось подключиться к базе: %w", err)
	}

	return pool, nil
}

func InitPool(ctx context.Context) {
	pool, err := newPool(ctx)
	if err != nil {
		log.Fatalf("Ошибка при инициализации базы: %v", err)
	}

	defer pool.Close()
	log.Println("Подключение к базе PostgreSQL установлено!")
}
