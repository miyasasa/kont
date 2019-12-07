package bitbucket

type Pagination struct {
	Size          int
	Limit         int
	IsLastPage    bool
	NextPageStart int
}

type PRPagination struct {
	Pagination
	Values []PullRequest
}

type UserPagination struct {
	Pagination
	Values []UserValues
}

func (p *UserPagination) GetUsers() []User {
	users := make([]User, 0)

	for _, v := range p.Values {
		users = append(users, v.User)
	}

	return users
}

type UserValues struct {
	User User `json:"user"`
}
