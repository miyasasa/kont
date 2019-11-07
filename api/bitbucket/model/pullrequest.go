package model

type PullRequest struct {
	Id          int32
	Version     int32
	Title       string
	Description string
	Reviewers   []Reviewer
}
