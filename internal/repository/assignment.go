package repository

import (
	"github.com/deckarep/golang-set"
	"kont/internal/common"
	"log"
	"math/rand"
	"time"
)

const (
	FIRST            = "FIRST"
	RANDOM           = "RANDOM"
	RANDOMINAVALABLE = "RANDOMINAVALABLE"
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

func (s *Stage) getRandomInAvailableReviewers(busyReviewers mapset.Set, owner common.Reviewer) common.Reviewer {
	log.Printf("BusyReviewers for per PR --> : %v and Owner--> %v", busyReviewers.ToSlice(), owner)
	return common.Reviewer{}
}

func (s *Stage) getRandomReviewer() common.Reviewer {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(s.Reviewers))
	return s.Reviewers[index]
}

func (s *Stage) getFirst() common.Reviewer {
	return s.Reviewers[0]
}
