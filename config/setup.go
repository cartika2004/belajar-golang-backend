package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	RedisAddr  string `mapstructure:"REDIS_ADDR"`
}

var ENV Config

func LoadConfig() {

	viper.AutomaticEnv()

	viper.AddConfigPath(".")      // Cari file di root folder
	viper.SetConfigName("config") // Nama filenya 'config'
	viper.SetConfigType("yaml")   // Tipenya YAML

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal("Error loading config: ", err)
	}

	log.Println(">>> Config Loaded Successfully using Viper <<<")
}
