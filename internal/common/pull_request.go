package common

type PR struct {
	Id          int32       `json:"id"`
	Version     int32       `json:"version"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Reviewers   []*Reviewer `json:"reviewers"`
	Author      Author      `json:"author"`
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

func (a Author) GetAuthorAsReviewer() Reviewer {
	return Reviewer{User: a.User}
}
