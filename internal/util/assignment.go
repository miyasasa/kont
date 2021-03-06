package util

import (
	"github.com/deckarep/golang-set"
	"kont/internal/common"
	"math/rand"
	"sort"
	"time"
)

func GetReviewerRandomly(reviewers mapset.Set) *common.Reviewer {
	cardinality := reviewers.Cardinality()
	if cardinality == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(cardinality)
	return reviewers.ToSlice()[index].(*common.Reviewer)
}

func GetFirstAvailableReviewerByPriority(reviewers mapset.Set) *common.Reviewer {
	if reviewers.Cardinality() == 0 {
		return nil
	}

	rev := reviewers.ToSlice()
	sort.Slice(rev, func(i, j int) bool {
		return rev[i].(*common.Reviewer).Priority > rev[j].(*common.Reviewer).Priority
	})

	return rev[0].(*common.Reviewer)
}
