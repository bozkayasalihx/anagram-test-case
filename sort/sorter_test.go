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
	tests := []structure{{Expected: "aabb", Actual: "baba"}, {Expected: "aenn", Actual: "anne"}}

	sorter := NewSortString()
	for _, test := range tests {
		assert.Equal(t, test.Expected, sorter(test.Actual))
	}
}
