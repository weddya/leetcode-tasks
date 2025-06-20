package valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("Opening and closing brackets of the same type", func(t *testing.T) {
		assert.Equal(t, true, isValid("()"))
	})

	t.Run("Opening and closing brackets of several types one by one", func(t *testing.T) {
		assert.Equal(t, true, isValid("()[]{}"))
	})

	t.Run("Opening and closing brackets of the same type", func(t *testing.T) {
		assert.Equal(t, false, isValid("(]"))
	})

	t.Run("Opening and closing brackets inside other brackets", func(t *testing.T) {
		assert.Equal(t, true, isValid("([])"))
	})
}

func TestIsValidTableDriven(t *testing.T) {
	testCases := []struct {
		Name          string
		InputValue    string
		ExpectedValue bool
	}{
		{
			Name:          "Opening and closing brackets of the same type",
			InputValue:    "()",
			ExpectedValue: true,
		},
		{
			Name:          "Opening and closing brackets of several types one by one",
			InputValue:    "()[]{}",
			ExpectedValue: true,
		},
		{
			Name:          "Opening and closing brackets of the same type",
			InputValue:    "(]",
			ExpectedValue: false,
		},
		{
			Name:          "Opening and closing brackets inside other brackets",
			InputValue:    "([])",
			ExpectedValue: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.EqualValues(t, tc.ExpectedValue, isValid(tc.InputValue))
		})
	}
}
