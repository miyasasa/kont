package common

type Reviewer struct {
	User User `json:"user"`
}

type User struct {
	Name string `json:"name"`
}
