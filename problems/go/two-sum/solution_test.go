package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("Test case #1", func(t *testing.T) {
		assert.Equal(t, []int{0, 1}, TwoSum([]int{2, 7, 11, 15}, 9))
	})
	t.Run("Test case #2", func(t *testing.T) {
		assert.Equal(t, []int{1, 2}, TwoSum([]int{3, 2, 4}, 6))
	})
	t.Run("Test case #3", func(t *testing.T) {
		assert.Equal(t, []int{0, 1}, TwoSum([]int{3, 3}, 6))
	})
}

func TestTwoSumTableDriven(t *testing.T) {
	testCases := []struct {
		Name          string
		InputSlice    []int
		InputTarget   int
		ExpectedSlice []int
	}{
		{
			Name:          "Test case #1",
			InputSlice:    []int{2, 7, 11, 15},
			InputTarget:   9,
			ExpectedSlice: []int{0, 1},
		},
		{
			Name:          "Test case #2",
			InputSlice:    []int{3, 2, 4},
			InputTarget:   6,
			ExpectedSlice: []int{1, 2},
		},
		{
			Name:          "Test case #3",
			InputSlice:    []int{3, 3},
			InputTarget:   6,
			ExpectedSlice: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.EqualValues(t, tc.ExpectedSlice, TwoSum(tc.InputSlice, tc.InputTarget))
		})
	}
}
