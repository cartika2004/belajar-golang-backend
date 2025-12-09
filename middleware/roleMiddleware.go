package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"project-todo/database"
)

// Satpam Khusus Admin
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil UserID yang sudah ditempel sama AuthMiddleware sebelumnya
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda belum login!"})
			c.Abort()
			return
		}

		// 2. Cek ke Redis: "Orang dengan ID ini, role-nya apa?"
		redisKey := fmt.Sprintf("user_role:%d", userID)
		role, err := database.Rdb.Get(database.Ctx, redisKey).Result()

		if err != nil {
			// Kalau data gak ada di Redis (misal expired), tolak
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi habis atau tidak valid (Redis)"})
			c.Abort()
			return
		}

		// 3. Cek Logic: Admin Only
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "MAAF! Halaman ini khusus ADMIN. Anda cuma User biasa."})
			c.Abort()
			return
		}

		// Kalau lolos, silakan lewat
		c.Next()
	}
}