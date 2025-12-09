package database

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"project-todo/config"
)

// Variabel Global biar bisa dipanggil di mana-mana
var Rdb *redis.Client
var Ctx = context.Background() // Redis butuh "Context" (aturan main Golang jaman now)

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.ENV.RedisAddr, // Ambil dari Viper
		Password: "",                   // Password kosong (default docker)
		DB:       0,                    // Default DB
	})

	// Test koneksi (Ping)
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Gagal konek ke Redis: ", err)
	}

	fmt.Println(">>> Sukses Konek ke Redis! (Si Kilat) <<<")
}