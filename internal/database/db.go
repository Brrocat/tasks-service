package database

import (
	"github.com/Brrocat/tasks-service/internal/task"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Bogdan_20 dbname=tasks port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		return nil, err
	}

	log.Println("Database connection successful")

	err = db.AutoMigrate(&task.Task{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
		return nil, err
	}

	return db, nil
}
