package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"project-todo/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbPassword := os.Getenv("DB_PASSWORD")
	
	dsn := fmt.Sprintf("sqlserver://sa:%s@localhost:1433?database=todo_app", dbPassword)
	
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}

	database.AutoMigrate(&models.User{}, &models.Todo{})

	DB = database
	log.Println(">>> Database Connected & Migrated <<<")
}