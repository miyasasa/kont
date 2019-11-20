package bitbucket

import (
	"kont/internal/common"
	"kont/internal/repository"
	"log"
)

func Listen(repo *repository.Repository) {
	log.Println("Bitbucket-PR is listening....")
	prs := fetchPRs(repo, 0)

	newPrs := filterPullRequestsHasNotReviewer(prs)
	log.Printf("LatestPRCount: %v", len(newPrs))

	repo.PRs = newPrs
	repo.Assign()

	updatePRs(repo)
}

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

func filterPullRequestsHasNotReviewer(prList []common.PullRequest) []common.PullRequest {
	prs := make([]common.PullRequest, 0)

	for _, v := range prList {
		if !v.DoesHaveAnyReviewer() {
			prs = append(prs, v)
		}
	}

	return prs
}
