package controllers

import (
	"fmt"
	"net/http"
	"time"

	"project-todo/config"
	"project-todo/database"
	"project-todo/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal daftar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil daftar!"})
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Salah username/password"})
		return
	}
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Salah username/password"})
		return
	}

	// token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := []byte(config.ENV.JWTSecret)
	tokenString, _ := token.SignedString(jwtSecret)

	redisKey := fmt.Sprintf("user_role:%d", user.ID)
	err := database.Rdb.Set(database.Ctx, redisKey, user.Role, 24*time.Hour).Err()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan ke Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"role":  user.Role, // Kasih tau di response dia login sebagai apa
	})
}