package common

type Reviewer struct {
	User     User `json:"user"`
	Approved bool `json:"approved"`
}

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}
