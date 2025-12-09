package main

import (
	"github.com/gin-gonic/gin"
	
	"project-todo/config"      // Import folder config
	"project-todo/database"
	"project-todo/controllers"
	"project-todo/middleware"
)

func main() {
	// 1. Load Config pakai Viper
	config.LoadConfig()

	// 2. Baru konek Database
	database.ConnectDatabase()

	database.ConnectRedis()

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

		adminGroup := protected.Group("/admin")
		adminGroup.Use(middleware.AdminOnly())

		{
			adminGroup.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Halo Bos Admin! Ini data rahasia."})
			})
	}}

	r.Run(":8080")
}