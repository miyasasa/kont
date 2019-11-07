package model

import (
	"encoding/json"
	"log"
)

type Pagination struct {
	Size       int
	Limit      int
	IsLastPage bool
	Values     []PullRequest
}

func (page *Pagination) ToString() string {
	s, e := json.Marshal(page)

	if e != nil {
		log.Println("Pagination::ToString, Page can not converted json")
	}

	return string(s)
}
