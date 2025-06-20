package palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("Correct palindrome num", func(t *testing.T) {
		assert.Equal(t, true, isPalindrome(121))
	})
	t.Run("Negative palindrome num", func(t *testing.T) {
		assert.Equal(t, false, isPalindrome(-121))
	})
	t.Run("Not palindrome num", func(t *testing.T) {
		assert.Equal(t, false, isPalindrome(10))
	})
}

func TestIsPalindromeTableDriven(t *testing.T) {
	testCases := []struct {
		Name          string
		InputNum      int
		ExpectedValue bool
	}{
		{
			Name:          "Correct palindrome num",
			InputNum:      121,
			ExpectedValue: true,
		},
		{
			Name:          "Negative palindrome num",
			InputNum:      -121,
			ExpectedValue: false,
		},
		{
			Name:          "Not palindrome num",
			InputNum:      10,
			ExpectedValue: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.EqualValues(t, tc.ExpectedValue, isPalindrome(tc.InputNum))
		})
	}
}
