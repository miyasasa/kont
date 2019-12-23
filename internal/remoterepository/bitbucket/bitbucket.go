package bitbucket

import (
	"kont/internal/common"
	"kont/internal/repository"
	"log"
)

func Listen(repo *repository.Repository) {
	log.Printf("Repo: %s --> is listening ...", repo.Name)
	fetchPRs(repo, 0)

	repo.AssignReviewersToPrs()

	updatePRs(repo)
}

// refactor with set // concat two two slice without duplicates
func UpdateUsers(repo *repository.Repository) {
	projectUsers := fetchUsers(repo.FetchProjectUsersUrl, repo.Token, 0)
	repoUsers := fetchUsers(repo.FetchRepoUsersUrl, repo.Token, 0)

	projectUsers = append(projectUsers, repoUsers...)

	users := make(map[string]common.User, 0)
	for _, u := range projectUsers {
		users[u.Name] = u
	}

	repo.Users = users
}
