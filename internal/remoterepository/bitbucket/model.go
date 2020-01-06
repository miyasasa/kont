package bitbucket

import (
	"kont/internal/common"
	"strings"
)

type PullRequestDefaultCommentUpdateModel struct {
	Text string `json:"text"`
}

func NewPullRequestDefaultCommentUpdateModel(text string, author string) PullRequestDefaultCommentUpdateModel {
	text = strings.Replace(text, "{{name}}", author, -1)
	return PullRequestDefaultCommentUpdateModel{Text: text}
}

type PullRequestReviewersUpdateModel struct {
	Id          int32      `json:"id"`
	Version     int32      `json:"version"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Reviewers   []Reviewer `json:"reviewers"`
}

type Reviewer struct {
	User User `json:"user"`
}

type User struct {
	Name string `json:"name"`
}

func MapPullRequestToUpdateModel(pr common.PullRequest) PullRequestReviewersUpdateModel {
	uPr := PullRequestReviewersUpdateModel{}
	uPr.Id = pr.Id
	uPr.Version = pr.Version
	uPr.Title = pr.Title
	uPr.Description = pr.Description
	uPr.Reviewers = mapReviewers(pr.Reviewers)

	return uPr
}

func mapReviewers(reviewers []*common.Reviewer) []Reviewer {
	var uRvs []Reviewer
	for _, r := range reviewers {
		uRv := Reviewer{User: User{Name: r.User.Name}}
		uRvs = append(uRvs, uRv)
	}

	return uRvs
}
