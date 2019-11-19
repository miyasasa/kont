package bitbucket

import (
	"log"
	"miya/internal/common"
	"miya/internal/repository"
)

func Listen(repo *repository.Repository) {
	log.Println("Bitbucket-PR is listening....")
	fetchPRs(repo)

	filterToGetLatestPullRequests(repo)
	log.Printf("LatestPRCount: %d", len(repo.PRs))

	repo.Assign()

	updatePRs(repo)
}

func UpdateUsers(repo *repository.Repository) {
	users := make(map[string]common.User, 0)

	projectUsers := fetchProjectUsers(repo)
	repoUsers := fetchRepositoryUsers(repo)

	projectUsers = append(projectUsers, repoUsers...)

	for _, u := range projectUsers {
		users[u.Name] = u
	}

	repo.Users = users
}

// An array of pull requests have not have any reviewer
func filterToGetLatestPullRequests(repo *repository.Repository) {
	prs := make([]common.PullRequest, 0)

	for _, v := range repo.PRs {
		if !v.DoesHaveAnyReviewer() {
			prs = append(prs, v)
		}
	}

	repo.PRs = prs
}
