package palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert.True(t, isPalindrome(121))
	assert.False(t, isPalindrome(-121))
	assert.False(t, isPalindrome(10))
}
