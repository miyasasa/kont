package repository

import (
	"github.com/deckarep/golang-set"
	"kont/init/env"
	"kont/internal/common"
	"log"
)

const (
	GITHUB    = "GITHUB"
	BITBUCKET = "BITBUCKET"
	GITLAB    = "GITLAB"
)

type Repository struct {
	FetchRepoUsersUrl    string                 `json:"fetchRepoUsersUrl"`
	FetchProjectUsersUrl string                 `json:"fetchProjectUsersUrl"`
	FetchPrsUrl          string                 `json:"fetchPrsUrl"`
	Token                string                 `json:"token"`
	ProjectName          string                 `json:"projectName"`
	Name                 string                 `json:"name"`
	Provider             string                 `json:"provider"`
	Users                map[string]common.User `json:"users"`
	Stages               []Stage                `json:"stages"`
	PRs                  []common.PullRequest   `json:"prs"`
}

func (repo *Repository) Initialize() {
	// choose according provider Bitbucket
	repo.FetchRepoUsersUrl = env.BitbucketFetchRepoUsersURL(repo.ProjectName, repo.Name)
	repo.FetchProjectUsersUrl = env.BitbucketFetchProjectUsersURL(repo.ProjectName)
	repo.FetchPrsUrl = env.BitbucketFetchPrListURL(repo.ProjectName, repo.Name)
}

func (repo *Repository) AssignReviewersToPrs() {

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	repo.filterPullRequestsHasNotReviewer()

	log.Printf("LatestPRCount: %v", len(repo.PRs))

	for i, pr := range repo.PRs {
		ownerAndReviewers := mapset.NewSet(repo.findUserInReviewers(pr.Author.User))
		for j, s := range repo.Stages {
			reviewer := s.GetReviewer(busyReviewers, ownerAndReviewers)
			if reviewer == nil {
				log.Printf("Reviewer not found in the stage")
				reviewer = repo.getReviewer(j, busyReviewers, ownerAndReviewers)
			}

			if reviewer != nil {
				ownerAndReviewers.Add(reviewer)
				repo.PRs[i].Reviewers = append(repo.PRs[i].Reviewers, reviewer)
			}
		}
	}
}

func (repo *Repository) getReviewer(index int, busyReviewers mapset.Set, ownerAndReviewers mapset.Set) *common.Reviewer {

	stages := append(repo.Stages[index+1:], repo.Stages[0:index]...)
	for _, s := range stages {
		reviewer := s.GetReviewer(busyReviewers, ownerAndReviewers)

		if reviewer != nil {
			return reviewer
		}
	}

	return nil
}

func (repo *Repository) filterPullRequestsHasNotReviewer() {
	prs := make([]common.PullRequest, 0)

	for _, v := range repo.PRs {
		if !v.IsAssignedAnyReviewer() {
			prs = append(prs, v)
		}
	}

	repo.PRs = prs
}

// which reviewers are busy :)
func (repo *Repository) getAssignedAndDoesNotApproveReviewers() mapset.Set {
	reviewers := mapset.NewSet()
	for _, pr := range repo.PRs {
		reviewers = reviewers.Union(pr.GetReviewersByUnApproved())
	}

	return reviewers
}

func (repo *Repository) findUserInReviewers(user common.User) *common.Reviewer {

	for _, s := range repo.Stages {
		if reviewer := s.getReviewerByUser(user); reviewer != nil {
			return reviewer
		}
	}

	return nil
}
