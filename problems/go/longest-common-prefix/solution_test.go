package longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("Has prefix", func(t *testing.T) {
		assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	})
	t.Run("Hasn't prefix", func(t *testing.T) {
		assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	})
}

func TestLongestCommonPrefixTableDriven(t *testing.T) {
	testCases := []struct {
		Name          string
		InputValue    []string
		ExpectedValue string
	}{
		{
			Name:          "Has prefix",
			InputValue:    []string{"flower", "flow", "flight"},
			ExpectedValue: "fl",
		},
		{
			Name:          "Hasn't prefix",
			InputValue:    []string{"dog", "racecar", "car"},
			ExpectedValue: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.EqualValues(t, tc.ExpectedValue, longestCommonPrefix(tc.InputValue))
		})
	}
}
