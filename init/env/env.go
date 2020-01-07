package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

var ServerPort = getEnv("SERVER_PORT", "1903")
var doOnce sync.Once

func getEnv(key string, defaultVal string) string {
	loadEnv()

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func loadEnv() {
	doOnce.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
		}
	})
}
