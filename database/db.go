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
	dbHost := config.ENV.DBHost
	if dbHost == "" {
		dbHost = "localhost"
	}
	
	dsn := fmt.Sprintf("sqlserver://sa:%s@%s:1433?database=todo_app", dbPassword, dbHost)
	
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}

	database.AutoMigrate(&models.User{}, &models.Todo{})

	DB = database
	log.Println(">>> Database Connected & Migrated <<<")
}