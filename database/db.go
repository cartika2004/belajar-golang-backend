package database

import (
	"fmt"
	"log"

	"project-todo/config"
	"project-todo/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbPassword := config.ENV.DBPassword
	
	dsn := fmt.Sprintf("sqlserver://sa:%s@localhost:1433?database=todo_app", dbPassword)
	
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}

	database.AutoMigrate(&models.User{}, &models.Todo{})

	DB = database
	log.Println(">>> Database Connected & Migrated <<<")
}