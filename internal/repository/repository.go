package repository

import (
	"github.com/deckarep/golang-set"
	"kont/internal/common"
	"kont/internal/util"
	"log"
)

const (
	GITHUB    = "GITHUB"
	BITBUCKET = "BITBUCKET"
	GITLAB    = "GITLAB"
)

type Repository struct {
	Host                 string                 `json:"host"`
	FetchRepoUsersUrl    string                 `json:"fetchRepoUsersUrl"`
	FetchProjectUsersUrl string                 `json:"fetchProjectUsersUrl"`
	FetchPrsUrl          string                 `json:"fetchPrsUrl"`
	Token                string                 `json:"token"`
	ProjectName          string                 `json:"projectName"`
	Name                 string                 `json:"name"`
	DevelopmentBranch    string                 `json:"developmentBranch"`
	Provider             string                 `json:"provider"`
	Users                map[string]common.User `json:"users"`
	Stages               []Stage                `json:"stages"`
	PRs                  []common.PullRequest   `json:"prs"`
}

func (repo *Repository) Initialize() {
	// choose according provider Bitbucket
	repo.FetchRepoUsersUrl = util.BitbucketFetchRepoUsersURL(repo.Host, repo.ProjectName, repo.Name)
	repo.FetchProjectUsersUrl = util.BitbucketFetchProjectUsersURL(repo.Host, repo.ProjectName)
	repo.FetchPrsUrl = util.BitbucketFetchPrListURL(repo.Host, repo.ProjectName, repo.Name)
}

func (repo *Repository) AssignReviewersToPrs() {

	busyReviewers := repo.getAssignedAndDoesNotApproveReviewers()

	repo.filterPullRequestByDevelopmentBranch()
	repo.filterPullRequestsHasNotReviewer()

	log.Printf("Repo: %s --> LatestPRCount: %v", repo.Name, len(repo.PRs))

	for i, pr := range repo.PRs {
		ownerAndReviewers := mapset.NewSet(repo.findReviewerByUsernameStage(pr.Author.User.Name))

		for _, s := range repo.Stages {
			reviewer := s.GetReviewer(busyReviewers, ownerAndReviewers)

			if reviewer == nil {
				reviewer = repo.getReviewer(i, busyReviewers, ownerAndReviewers)
			}

			if reviewer != nil {
				ownerAndReviewers.Add(reviewer)
				repo.PRs[i].Reviewers = append(repo.PRs[i].Reviewers, reviewer)
			}
		}
	}
}

func (repo *Repository) getReviewer(index int, busyReviewers mapset.Set, ownerAndReviewers mapset.Set) *common.Reviewer {

	stages := append(repo.Stages[index:], repo.Stages[0:index]...)
	for _, s := range stages {
		reviewer := s.GetReviewer(busyReviewers, ownerAndReviewers)
		if reviewer != nil {
			return reviewer
		}
	}
	return nil
}

func (repo *Repository) filterPullRequestByDevelopmentBranch() {
	prs := make([]common.PullRequest, 0)

	for _, p := range repo.PRs {
		if repo.DevelopmentBranch == p.ToRef.DisplayId {
			prs = append(prs, p)
		}
	}

	repo.PRs = prs
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
		reviewers = reviewers.Union(repo.GetReviewersByUnApproved(pr))
	}

	return reviewers
}

func (repo *Repository) GetReviewersByUnApproved(pr common.PullRequest) mapset.Set {
	reviewers := mapset.NewSet()

	for _, r := range pr.Reviewers {
		if !r.Approved {
			rv := repo.findReviewerByUsernameStage(r.User.Name)
			if rv != nil {
				reviewers.Add(rv)
			}
		}
	}

	return reviewers
}

func (repo *Repository) findReviewerByUsernameStage(username string) *common.Reviewer {

	for _, s := range repo.Stages {
		if reviewer := s.getReviewerByUserName(username); reviewer != nil {
			return reviewer
		}
	}

	return nil
}
