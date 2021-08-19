package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	input := []int{1, 2, 5, 5, 4}
	nextPermutation(input)
	assert.EqualValues(t, input, []int{1, 4, 2, 5, 5})
}
