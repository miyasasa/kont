package repository

import (
	"kont/internal/common"
	"math/rand"
	"time"
)

const (
	FIRST  = "FIRST"
	RANDOM = "RANDOM"
)

type Stage struct {
	Name      string
	Reviewers []common.Reviewer
	Policy    string
	// added policy for assignment
}

func (s *Stage) GetReviewer() common.Reviewer {
	switch s.Policy {
	case FIRST:
		return s.getFirst()
	case RANDOM:
		return s.getRandomReviewer()
	default:
		return s.getRandomReviewer()
	}
}

func (s *Stage) getRandomReviewer() common.Reviewer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Reviewers))
	return s.Reviewers[index]
}

func (s *Stage) getFirst() common.Reviewer {
	return s.Reviewers[0]
}
