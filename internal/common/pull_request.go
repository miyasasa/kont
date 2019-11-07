package common

type PullRequest struct {
	Id          int32
	Version     int32
	Title       string
	Description string
	Reviewers   []Reviewer
}

func (pr *PullRequest) DoesHaveAnyReviewer() bool {
	return len(pr.Reviewers) != 0
}
