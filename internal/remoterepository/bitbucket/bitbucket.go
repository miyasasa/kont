package bitbucket

import (
	"kont/internal/client"
	"kont/internal/common"
	"kont/internal/repository"
	"log"
)

var bitbucketApi = NewBitbucketApi(client.HttpClientInstance)

func Listen(repo *repository.Repository) {
	log.Printf("Repo: %s --> start to observe ...", repo.Name)

	bitbucketApi.fetchPRs(repo, 0)

	repo.AssignReviewersToPrs()

	bitbucketApi.updatePRs(repo)
}

func UpdateUsers(repo *repository.Repository) {
	projectUsers := bitbucketApi.fetchUsers(repo.FetchProjectUsersUrl, repo.Token, 0)
	repoUsers := bitbucketApi.fetchUsers(repo.FetchRepoUsersUrl, repo.Token, 0)

	projectUsers = append(projectUsers, repoUsers...)

	users := make(map[string]common.User, 0)
	for _, u := range projectUsers {
		users[u.Name] = u
	}

	repo.Users = users
}
