package common

import "github.com/deckarep/golang-set"

type PullRequest struct {
	Id          int32      `json:"id"`
	Version     int32      `json:"version"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Reviewers   []Reviewer `json:"reviewers"`
	Author      Author     `json:"author"`
}

type Reviewer struct {
	User     User `json:"user"`
	Approved bool `json:"approved"`
}

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type Author struct {
	User User `json:"user"`
}

func (a Author) GetAuthorAsReviewer() Reviewer {
	return Reviewer{User: a.User}
}

func (pr *PullRequest) IsAssignedAnyReviewer() bool {
	return len(pr.Reviewers) != 0
}

func (pr *PullRequest) GetReviewersByUnApproved() mapset.Set {
	reviewers := mapset.NewSet()

	for _, r := range pr.Reviewers {
		if !r.Approved {
			reviewers.Add(r)
		}
	}

	return reviewers
}
