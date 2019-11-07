package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var BitbucketFetchPrListUrl = getEnv("BITBUCKET_FETCH_PR_LIST_URL", "")
var BitbucketToken = getEnv("BITBUCKET_TOKEN", "")

func getEnv(key string, defaultVal string) string {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	return getEnvFromOS(key, defaultVal)
}

func getEnvFromOS(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
