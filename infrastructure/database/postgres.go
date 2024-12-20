package database

import (
	"fmt"
	"log"

	"github.com/gabrielmoura33/neoway-backend-test/config"
	"github.com/gabrielmoura33/neoway-backend-test/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&domain.Client{}); err != nil {
		log.Panicln("Failed to migrate database:", err)
	}

	return db
}
