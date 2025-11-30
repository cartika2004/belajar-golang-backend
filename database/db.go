package database

import (
	"fmt"
	"log"
	"os" // Buat baca environment variable

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"project-todo/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ambil password dari .env
	dbPassword := os.Getenv("DB_PASSWORD")
	
	// Masukkan ke dalam string koneksi pakai fmt.Sprintf
	dsn := fmt.Sprintf("sqlserver://sa:%s@localhost:1433?database=todo_app", dbPassword)
	
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}

	database.AutoMigrate(&models.User{}, &models.Todo{})

	DB = database
	log.Println(">>> Database Connected & Migrated <<<")
}