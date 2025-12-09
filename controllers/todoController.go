package controllers

import (
	"net/http"

	"project-todo/database"
	"project-todo/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Deadline == " " {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Deadline jangan kosong cuy!"})
	return}
	
	userID, _ := c.Get("userID")
	input.UserID = userID.(uint)

	database.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"message": "Todo dibuat!", "data": input})
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	userID, _ := c.Get("userID")
	database.DB.Where("user_id = ?", userID).Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")
	var todo models.Todo

	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&todo).Updates(input)
	c.JSON(http.StatusOK, gin.H{"message": "Updated!", "data": todo})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")
	var todo models.Todo
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted!"})
}
