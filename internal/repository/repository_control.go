package repository

import (
	"miya/init/env"
	"miya/internal/common"
)

// get all repositories
func GetAllRepositories() []Repository {
	return []Repository{getRepository()}
}

func getRepository() Repository {
	repo := new(Repository)
	repo.FetchPrURL = env.BitbucketFetchPrListUrl
	repo.Token = env.BitbucketToken
	repo.Provider = BITBUCKET
	repo.ProjectName = "BESG"
	repo.Name = "core-network"
	repo.Reviewers = getReviewers()

	return *repo
}

//it will initial from remote-repository not manually
func getReviewers() map[string][]common.Reviewer {
	rv := make(map[string][]common.Reviewer)
	rv["STAGE1"] = []common.Reviewer{getReviewer("ataday"), getReviewer("baydogdu"), getReviewer("huseyiny"), getReviewer("eunal")}
	rv["STAGE2"] = []common.Reviewer{getReviewer("veroglu")}
	rv["STAGE3"] = []common.Reviewer{getReviewer("ykizilkaya")}

	return rv
}

func getReviewer(name string) common.Reviewer {
	return common.Reviewer{User: common.User{Name: name}}
}
