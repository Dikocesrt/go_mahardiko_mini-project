package configs

import (
	"habit/repositories/mysql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitConfigMySQL() mysql.Config {
	return mysql.Config{
		DBName: os.Getenv("DBName"),
		DBUser: os.Getenv("DBUser"),
		DBPass: os.Getenv("DBPass"),
		DBHost: os.Getenv("DBHost"),
		DBPort: os.Getenv("DBPort"),
		DBUnix: os.Getenv("DBUnix"),
	}
}

func InitConfigCloudinary() string {
	return os.Getenv("CLOUDINARY_URL")
}

func InitConfigKeyChatbot() string {
	return os.Getenv("KEY")
}