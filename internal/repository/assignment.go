package repository

import (
	"github.com/deckarep/golang-set"
	"kont/internal/common"
	"kont/internal/util"
)

const (
	BYPRIORITYINAVAILABLE = "BYPRIORITYINAVAILABLE"
	RANDOMINAVAILABLE     = "RANDOMINAVAILABLE"
)

type Stage struct {
	Name      string
	Reviewers []*common.Reviewer
	Policy    string
}

// For getting Difference over map-set, items have to same pointer-address to compare
func (s *Stage) GetReviewer(busyReviewers mapset.Set, ownerAndReviewers mapset.Set) *common.Reviewer {
	stageReviewers := mapset.NewSetFromSlice(s.getReviewers())
	availableReviewers := stageReviewers.Difference(busyReviewers).Difference(ownerAndReviewers)

	if availableReviewers.Cardinality() == 0 {
		reviewers := stageReviewers.Difference(ownerAndReviewers)
		if reviewers.Cardinality() == 0 {
			return nil // Reviewers are assigned already PR's reviewer-or-owner
		}

		return util.GetReviewerRandomly(reviewers)
	}

	if s.Policy == BYPRIORITYINAVAILABLE {
		return util.GetFirstAvailableReviewerByPriority(availableReviewers)
	}

	return util.GetReviewerRandomly(availableReviewers)
}

func (s *Stage) getReviewers() []interface{} {
	rv := make([]interface{}, len(s.Reviewers))
	for i, v := range s.Reviewers {
		rv[i] = v
	}

	return rv
}

func (s *Stage) getReviewerByUserName(username string) *common.Reviewer {
	for _, r := range s.Reviewers {
		if username == r.User.Name {
			return r
		}
	}

	return nil
}
