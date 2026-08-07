package main

import (
	"FinanceApp/internal/repository"
	"FinanceApp/internal/service"
	"FinanceApp/internal/transport"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println(".env файл не найден")
	}

	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("HTTP_PORT")

	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal("Ошибка создания пула соединений:", err)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(ctx); err != nil {
		log.Fatal("База данных недоступна:", err)
	}

	log.Println("Успешное подключение к БД")

	repo := repository.NewRepository(dbpool)

	userService := service.NewUserService(repo)
	userHandler := transport.NewUserHandler(userService)

	transactionService := service.NewTransactionService(repo)
	transactionHandler := transport.NewTransactionHandler(transactionService)

	router := transport.NewRouter(userHandler, transactionHandler)

	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(router.Start(port))
}
