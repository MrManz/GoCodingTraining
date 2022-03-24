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

func TestPermutationOfPalindrome(t *testing.T) {
	assert.Equal(t, false, isPermutationOfPalindrome("ABCA"))
	assert.Equal(t, true, isPermutationOfPalindrome("ABBA"))
	assert.Equal(t, true, isPermutationOfPalindrome("Taco cat"))
	assert.Equal(t, false, isPermutationOfPalindrome("aco cat"))
	assert.Equal(t, false, isPermutationOfPalindrome("Cat"))
	assert.Equal(t, false, isPermutationOfPalindrome("ADSA"))
}

func TestOneEditDifference(t *testing.T) {
	assert.Equal(t, true, isOneEditDifference("pale", "pale"))
	assert.Equal(t, true, isOneEditDifference("pale", "ple"))
	assert.Equal(t, true, isOneEditDifference("pales", "pale"))
	assert.Equal(t, true, isOneEditDifference("pale", "pales"))
	assert.Equal(t, true, isOneEditDifference("pale", "bale"))
	assert.Equal(t, false, isOneEditDifference("pale", "bake"))
}

func TestCompressString(t *testing.T) {
	assert.Equal(t, "a2b1c5a3", compressString("aabcccccaaa"))
	assert.Equal(t, "a1b1c5a3", compressString("abcccccaaa"))
}

func TestReverseMatrix(t *testing.T) {
	assert.Equal(t, [][]int{{7, 8, 9}, {4, 5, 6}, {1, 2, 3}}, reverseMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	assert.Equal(t, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, rotateMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
