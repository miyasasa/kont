package repository

import (
	"github.com/deckarep/golang-set"
	"kont/internal/common"
	"kont/internal/util"
)

const (
	BYORDERINAVAILABLE = "BYORDERINAVAILABLE"
	RANDOMINAVAILABLE  = "RANDOMINAVAILABLE"
)

type Stage struct {
	Name      string
	Reviewers []*common.Reviewer
	Policy    string
}

func (s *Stage) GetReviewer(busyReviewers mapset.Set, ownerAndReviewers mapset.Set) *common.Reviewer {
	availableReviewers := mapset.NewSetFromSlice(s.getReviewers()).Difference(busyReviewers).Difference(ownerAndReviewers)

	if availableReviewers.Cardinality() == 0 {
		reviewers := busyReviewers.Difference(ownerAndReviewers)
		if reviewers.Cardinality() == 0 {
			return nil // Reviewers are assigned already PR's reviewer-or-owner
		}

		return util.GetReviewerRandomly(reviewers)
	}

	if s.Policy == BYORDERINAVAILABLE {
		return util.GetFirstAvailableReviewerByOrder(availableReviewers)
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
