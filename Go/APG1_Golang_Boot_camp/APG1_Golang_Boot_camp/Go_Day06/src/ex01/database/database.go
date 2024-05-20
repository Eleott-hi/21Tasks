package database

import (
	"ex01/config"
	"ex01/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func New() *gorm.DB {
	if database == nil {
		database = create_connection()
	}

	return database
}

func create_connection() *gorm.DB {
	db_config := config.Config.DatabaseConfig

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		db_config.Host,
		db_config.User,
		db_config.Password,
		db_config.Name,
		db_config.Port,
		db_config.SSLMode,
		db_config.TimeZone,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&models.Article{})
	return database
}
