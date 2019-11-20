package common

type PRPagination struct {
	Size       int
	Limit      int
	IsLastPage bool
	Values     []PullRequest
}

type UserPagination struct {
	Size          int
	Limit         int
	IsLastPage    bool
	Values        []UserValues
	NextPageStart int
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
