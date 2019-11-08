package common

type PullRequest struct {
	Id          int32      `json:"id"`
	Version     int32      `json:"version"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Reviewers   []Reviewer `json:"reviewers"`
}

func (pr *PullRequest) DoesHaveAnyReviewer() bool {
	return len(pr.Reviewers) != 0
}
