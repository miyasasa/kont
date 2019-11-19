package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

var BitbucketToken = getEnv("BITBUCKET_TOKEN")
var doOnce sync.Once

// get project name & repo name dynamically from repo
func BitbucketFetchRepoUsersURL(projetName string, repoName string) string {
	return getEnv("BITBUCKET_BASE_URL") +
		getEnv("BITBUCKET_PROJECT_PATH") +
		projetName +
		getEnv("BITBUCKET_REPOSITORY_PATH") +
		repoName +
		getEnv("BITBUCKET_USER_PATH")

}

// get project name dynamically from repo
func BitbucketFetchProjectUsersURL(projetName string) string {
	return getEnv("BITBUCKET_BASE_URL") +
		getEnv("BITBUCKET_PROJECT_PATH") +
		projetName +
		getEnv("BITBUCKET_USER_PATH")
}

func BitbucketFetchPrListURL(projetName string, repoName string) string {
	return getEnv("BITBUCKET_BASE_URL") +
		getEnv("BITBUCKET_PROJECT_PATH") +
		projetName +
		getEnv("BITBUCKET_REPOSITORY_PATH") +
		repoName +
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
