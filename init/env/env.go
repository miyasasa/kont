package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

var BitbucketFetchRepositoryUsersURL = bitbucketFetchRepoUsersURL()
var BitbucketFetchProjectUsersURL = bitbucketFetchProjectUsersURL()
var BitbucketFetchPRURL = bitbucketFetchPrListURL()
var BitbucketToken = getEnv("BITBUCKET_TOKEN")

var doOnce sync.Once

// get project name & repo name dynamically from repo
func bitbucketFetchRepoUsersURL() string {
	return getEnv("BITBUCKET_BASE_URL") +
		getEnv("BITBUCKET_PROJECT_PATH") +
		getEnv("BITBUCKET_REPOSITORY_PATH") +
		getEnv("BITBUCKET_USER_PATH")

}

// get project name dynamically from repo
func bitbucketFetchProjectUsersURL() string {
	return getEnv("BITBUCKET_BASE_URL") +
		getEnv("BITBUCKET_PROJECT_PATH") +
		getEnv("BITBUCKET_USER_PATH")
}

func bitbucketFetchPrListURL() string {
	return getEnv("BITBUCKET_BASE_URL") +
		getEnv("BITBUCKET_PROJECT_PATH") +
		getEnv("BITBUCKET_REPOSITORY_PATH") +
		getEnv("BITBUCKET_PR_PATH")
}

func getEnv(key string) string {
	loadEnv()

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return ""
}

func loadEnv() {
	doOnce.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
		}
	})
}
