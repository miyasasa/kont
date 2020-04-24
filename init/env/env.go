package env

import (
	"github.com/joho/godotenv"
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
		_ = godotenv.Load()
	})
}
