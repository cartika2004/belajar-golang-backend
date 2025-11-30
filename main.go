package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Import library ini

	"project-todo/controllers"
	"project-todo/database"
	"project-todo/middleware"
)

func main() {
	// 1. Load file .env dulu!
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Baru konek Database
	database.ConnectDatabase()

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware)
	{
		protected.POST("/todos", controllers.CreateTodo)
		protected.GET("/todos", controllers.GetTodos)
		protected.PUT("/todos/:id", controllers.UpdateTodo)
		protected.DELETE("/todos/:id", controllers.DeleteTodo)
	}

	r.Run(":8080")
}