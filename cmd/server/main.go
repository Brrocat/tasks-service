package main

import (
	"log"

	"github.com/Brrocat/tasks-service/internal/database"
	"github.com/Brrocat/tasks-service/internal/task"
	transportgrpc "github.com/Brrocat/tasks-service/internal/transport/grpc"
)

func main() {
	// 1. Инициализация подключения к базе данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Database connection established")

	// 2. Инициализация репозитория и сервиса задач
	repo := task.NewRepository(db)
	svc := task.NewService(repo)
	log.Println("Task service initialized")

	// 3. Создание gRPC клиента для Users сервиса
	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051") // Адрес Users сервиса
	if err != nil {
		log.Fatalf("Failed to connect to Users service: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to Users service")

	// 4. Запуск gRPC сервера для Tasks сервиса
	log.Println("Starting Tasks gRPC server on :50052")
	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
