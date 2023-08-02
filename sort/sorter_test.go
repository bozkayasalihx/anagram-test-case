package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type structure struct {
	Expected string
	Actual   string
}

func TestSortString(t *testing.T) {
	tests := []structure{
		{Expected: "aabb", Actual: "baba"},
		{Expected: "aenn", Actual: "anne"},
		{Expected: "aaagmnr", Actual: "anagram"},
		{Expected: "below", Actual: "elbow"},
		{Expected: "dstuy", Actual: "study"},
		{Expected: "ghint", Actual: "night"},
		{Expected: "act", Actual: "cat"},
		{Expected: "deersst", Actual: "dessert"},
	}

	sorter := NewSortString()
	for _, test := range tests {
		assert.Equal(t, test.Expected, sorter(test.Actual))
	}
}
