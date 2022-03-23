package GoCodingTraining

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUnique(t *testing.T) {
	assert.Equal(t, true, hasUniqueCharacters("abcd"))
	assert.Equal(t, true, hasUniqueCharacters("a"))
	assert.Equal(t, false, hasUniqueCharacters("abca"))
	assert.Equal(t, false, hasUniqueCharacters("aabb"))
	assert.Equal(t, false, hasUniqueCharacters("cc"))
}

func TestIsPermutation(t *testing.T) {
	assert.Equal(t, true, isPermutation("ABC", "CBA"))
	assert.Equal(t, true, isPermutation("ABC", "BCA"))
	assert.Equal(t, true, isPermutation("ABC", "bca"))
	assert.Equal(t, false, isPermutation("ABC", "DCA"))
	assert.Equal(t, false, isPermutation("ABC", "AAA"))
	assert.Equal(t, false, isPermutation("ABC", "BAA"))
}

func TestReplaceSpacesInCharArray(t *testing.T) {
	assert.Equal(t, []rune("Mr%20John%20Smith"), replaceSpacesInCharArray([]rune("Mr John Smith")))
	assert.Equal(t, []rune("%20John%20Smith"), replaceSpacesInCharArray([]rune(" John Smith")))
	assert.Equal(t, []rune("MrJohnSmith"), replaceSpacesInCharArray([]rune("MrJohnSmith")))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, "ABCDE", reverse("EDCBA"))
	assert.Equal(t, "Cat", reverse("taC"))
}

func TestPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome("ABBA"))
	assert.Equal(t, true, isPalindrome("Taco cat"))
	assert.Equal(t, false, isPalindrome("aco cat"))
	assert.Equal(t, false, isPalindrome("Cat"))
	assert.Equal(t, false, isPalindrome("ADSA"))
}
