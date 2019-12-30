package common

type PullRequest struct {
	Id          int32       `json:"id"`
	Version     int32       `json:"version"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Reviewers   []*Reviewer `json:"reviewers"`
	Author      Author      `json:"author"`
	ToRef       ToRef       `json:"toRef"`
}

type Reviewer struct {
	Order    int  `json:"order"`
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

type ToRef struct {
	DisplayId string `json:"displayId"`
}

func (a Author) GetAuthorAsReviewer() Reviewer {
	return Reviewer{User: a.User}
}

func (pr *PullRequest) IsAssignedAnyReviewer() bool {
	return len(pr.Reviewers) != 0
}
