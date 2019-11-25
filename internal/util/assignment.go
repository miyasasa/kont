package util

import (
	"github.com/deckarep/golang-set"
	"kont/internal/common"
	"math/rand"
	"time"
)

func GetReviewerRandomly(set mapset.Set) *common.Reviewer {
	cardinality := set.Cardinality()
	if cardinality == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(cardinality)
	return set.ToSlice()[index].(*common.Reviewer)
}

func GetReviewerFirstAvailable(set mapset.Set) *common.Reviewer {
	if set.Cardinality() == 0 {
		return nil
	}

	return set.ToSlice()[0].(*common.Reviewer)
}
