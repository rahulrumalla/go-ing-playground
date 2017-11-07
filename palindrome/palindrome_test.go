package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrom(t *testing.T) {
	input := "HANNAH"
	assert.True(t, IsPalindrome(input), "Error")

	input = "NITIN"
	assert.True(t, IsPalindrome(input), "Error")

	input = "RAHUL"
	assert.False(t, IsPalindrome(input), "Error")
}
