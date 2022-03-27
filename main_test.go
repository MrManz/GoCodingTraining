package GoCodingTraining

import (
	"container/list"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
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

func TestBitFunctions(t *testing.T) {
	n := int64(31)

	fmt.Println(strconv.FormatInt(n, 2))
	fmt.Println(rune('A') - 64)
	fmt.Println(rune('Z') - 64)
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

func TestCamelCase(t *testing.T) {
	testStrings := []string{"S;M;makeCoffe()", "S;M;doStuff()", "S;C;OrangeHighlighter"}
	assert.Equal(t, []string{"make coffe", "do stuff", "orange highlighter"}, formatCamelCase(testStrings))
	testStrings = []string{"C;V;mobile phone", "C;C;coffee machine", "C;M;white sheet of paper"}
	assert.Equal(t, []string{"mobilePhone", "CoffeeMachine", "whiteSheetOfPaper()"}, formatCamelCase(testStrings))
}

func TestDivisibleSumPairs(t *testing.T) {
	var arr []int32 = []int32{43, 95, 51, 55, 40, 86, 65, 81, 51, 20, 47, 50, 65, 53, 23, 78, 75, 75, 47, 73, 25, 27, 14, 8, 26, 58, 95, 28, 3, 23, 48, 69, 26, 3, 73, 52, 34, 7, 40, 33, 56, 98, 71, 29, 70, 71, 28, 12, 18, 49, 19, 25, 2, 18, 15, 41, 51, 42, 46, 19, 98, 56, 54, 98, 72, 25, 16, 49, 34, 99, 48, 93, 64, 44, 50, 91, 44, 17, 63, 27, 3, 65, 75, 19, 68, 30, 43, 37, 72, 54, 82, 92, 37, 52, 72, 62, 3, 88, 82, 71}
	fmt.Println(len(arr))
	assert.Equal(t, int32(216), divisibleSumPairs(100, 22, arr))
}

func TestMatchingStrings(t *testing.T) {
	var arrStrings = []string{"aba", "baba", "aba", "xzxb"}
	var arrQueries = []string{"aba", "xzxb", "ab"}
	assert.Equal(t, []int32{2, 1, 0}, matchingStrings(arrStrings, arrQueries))
}

func TestLonelyInts(t *testing.T) {
	var arr = []int32{1, 1, 2, 5, 5}
	assert.Equal(t, int32(2), lonelyInteger(arr))
}

func TestGradeStudents(t *testing.T) {
	var arr = []int32{84, 29, 57}
	assert.Equal(t, []int32{85, 29, 57}, gradingStudents(arr))
}

func TestDiagonalDifference(t *testing.T) {
	assert.Equal(t, int32(2), diagonalDifference([][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}))
}

func TestReverseRow(t *testing.T) {
	assert.Equal(t, [][]int32{{9, 2, 3}, {4, 5, 6}, {1, 8, 9}}, reverseRow([][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}, 0))
	assert.Equal(t, [][]int32{{7, 2, 3, 3}, {2, 5, 6, 3}, {4, 8, 9, 3}, {1, 8, 9, 3}}, reverseRow([][]int32{{1, 2, 3, 3}, {4, 5, 6, 3}, {2, 8, 9, 3}, {7, 8, 9, 3}}, 0))
}

func TestSetMatrixRowsColumnsZero(t *testing.T) {
	assert.Equal(t, [][]int32{{0, 0, 0}, {4, 0, 6}, {9, 0, 9}}, setMatrixRowsColumnsZero([][]int32{{1, 0, 3}, {4, 5, 6}, {9, 8, 9}}))
	assert.Equal(t, [][]int32{{1, 2, 3, 0}, {0, 0, 0, 0}, {9, 8, 9, 0}}, setMatrixRowsColumnsZero([][]int32{{1, 2, 3, 4}, {4, 5, 6, 0}, {9, 8, 9, 7}}))
}

func TestIsRotatedString(t *testing.T) {
	assert.Equal(t, true, isRotatedString("waterbottle", "erbottlewat"))
	assert.Equal(t, false, isRotatedString("waterbottles", "erbottlewat"))

}

func TestDeleteDuplicatesFromList(t *testing.T) {
	l := list.New()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushFront(1)
	l.PushFront(1)

	l = deleteDuplicatesFromList(l)

	var values []int
	for e := l.Front(); e != nil; e = e.Next() {
		values = append(values, e.Value.(int))
	}
	assert.Equal(t, []int{2, 3, 1}, values)
}

func TestTwoArrays(t *testing.T) {
	var arr1 = []int32{2, 1, 3}
	var arr2 = []int32{7, 8, 9}
	assert.Equal(t, "YES", twoArrays(10, arr1, arr2))
	arr1 = []int32{1, 2, 2, 1}
	arr2 = []int32{3, 3, 3, 4}
	assert.Equal(t, "NO", twoArrays(10, arr1, arr2))
}

func TestHackerRankBirthday(t *testing.T) {
	assert.Equal(t, int32(2), hackerrankBirthday([]int32{1, 2, 1, 3, 2}, 3, 2))
}
