package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 定义配置文件里面的类型
type Config struct {
	AppPort string
	AppName string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
}

var App *Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	App = &Config{
		AppPort: getENV("APP_PORT", "8080"),
		AppName: getENV("APP_NAME", "MagicStreamMovieServer"),
		DBHost:  getENV("DB_HOST", "localhost"),
		DBPort:  getENV("DB_PORT", "27017"),
		DBUser:  getENV("DB_USER", ""),
		DBPass:  getENV("DB_PASS", ""),
		DBName:  getENV("DB_NAME", "magic-stream-movies"),
	}
}

// 定义一个默认的值 如果配置文件里面没有写
func getENV(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
