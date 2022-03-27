package GoCodingTraining

import (
	"container/list"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func reverse(input string) string {
	var returnValue string
	for _, value := range input {
		returnValue = string(value) + returnValue
	}
	return returnValue
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return val > 0
}

func setBit(n int, pos uint) int {
	n |= 1 << pos
	return n
}

func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

type sortableInt32array []int32

func (f sortableInt32array) Len() int {
	return len(f)
}

func (f sortableInt32array) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f sortableInt32array) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func twoArrays(k int32, A []int32, B []int32) string {
	var sortableA sortableInt32array = A
	var sortableB sortableInt32array = B

	sort.Sort(sortableA)
	sort.Sort(sortableB)

	for i, j := 0, len(sortableB)-1; i < j; i, j = i+1, j-1 {
		sortableB[i], sortableB[j] = sortableB[j], sortableB[i]
	}

	for i := range sortableA {
		if !(sortableA[i]+sortableB[i] >= k) {
			return "NO"
		}
	}
	return "YES"

}

func maximumPerimeterTriangle(sticks []int32) []int32 {
	maxPerimeter := int64(0)
	var returnArray sortableInt32array
	var sortableArray sortableInt32array = sticks

	sort.Sort(sortableArray)

	for i, j := 0, len(sortableArray)-1; i < j; i, j = i+1, j-1 {
		sortableArray[i], sortableArray[j] = sortableArray[j], sortableArray[i]
	}

	for i := 0; i < len(sortableArray)-2; i++ {
		if (sortableArray[i]+sortableArray[i+1] > sortableArray[i+2]) && (sortableArray[i]+sortableArray[i+2] > sortableArray[i+1]) && (sortableArray[i+1]+sortableArray[i+2] > sortableArray[i]) {
			if int64(sortableArray[i])+int64(sortableArray[i+1])+int64(sortableArray[i+2]) > int64(maxPerimeter) {
				returnArray = nil
				returnArray = append(returnArray, sortableArray[i], sortableArray[i+1], sortableArray[i+2])
				maxPerimeter = int64(sortableArray[i]) + int64(sortableArray[i+1]) + int64(sortableArray[i+2])
			}
		}
	}

	if maxPerimeter == int64(0) {
		return []int32{-1}
	} else {
		sort.Sort(returnArray)
		return returnArray
	}

}

func hackerrankBirthday(s []int32, d int32, m int32) int32 {
	segmentCounter := 0
	for i := 0; i < len(s)-int(m)+1; i++ {
		sumSegment := s[i]
		for offset := 1; offset < int(m); offset++ {
			sumSegment += s[i+offset]
		}
		if sumSegment == d {
			segmentCounter++
		}
	}
	return int32(segmentCounter)
}

func migratoryBirds(arr []int32) int32 {
	var m = make(map[int32]int32)

	for _, value := range arr {
		m[value]++
	}

	highestCount := int32(0)
	var highestKeys sortableInt32array
	for key, value := range m {
		if value > highestCount {
			highestKeys = nil
			highestKeys = append(highestKeys, key)
			highestCount = value
		} else if value == highestCount {
			highestKeys = append(highestKeys, key)
		}
	}
	sort.Sort(highestKeys)
	return highestKeys[0]

}

func setRowToZero(matrix [][]int32, index int) [][]int32 {
	for j := range matrix[index] {
		matrix[index][j] = 0
	}
	return matrix
}

func setColumnToZero(matrix [][]int32, j int) [][]int32 {
	for i := range matrix {
		matrix[i][j] = 0
	}
	return matrix
}

func isRotatedString(input1, input2 string) bool {
	return strings.Contains(input1+input1, input2)
}

func deleteDuplicatesFromList(l *list.List) *list.List {
	bitMap := make(map[any]bool)
	returnList := list.New()

	for e := l.Front(); e != nil; e = e.Next() {
		if !bitMap[e.Value] {
			returnList.PushFront(e.Value)
			bitMap[e.Value] = true
		}
	}

	return returnList
}

func setMatrixRowsColumnsZero(matrix [][]int32) [][]int32 {
	rowZero := make([]bool, len(matrix))
	columnZero := make([]bool, len(matrix[0]))

	for i, row := range matrix {
		for j, value := range row {
			if value == 0 {
				rowZero[i] = true
				columnZero[j] = true
			}
		}
	}

	for i := range matrix {
		if rowZero[i] {
			matrix = setRowToZero(matrix, i)
		}
	}

	for j := range matrix[0] {
		if columnZero[j] {
			matrix = setColumnToZero(matrix, j)
		}
	}

	return matrix
}

func reverseRow(matrix [][]int32, columnIndex int) [][]int32 {
	var returnMatrix = make([][]int32, len(matrix))
	for i := range returnMatrix {
		returnMatrix[i] = make([]int32, len(matrix))
	}

	for i, row := range matrix {
		for j, value := range row {
			if columnIndex == j {
				returnMatrix[i][j] = matrix[len(matrix)-1-i][j]
			} else {
				returnMatrix[i][j] = value
			}
		}
	}
	return returnMatrix
}

func diagonalDifference(arr [][]int32) int32 {
	var forwardSum int32 = 0
	var backwardSum int32 = 0

	for i, line := range arr {
		for j, value := range line {
			if i == j {
				forwardSum += value
			}
			if i+j == len(arr)-1 {
				backwardSum += value
			}
		}
	}

	return int32(math.Abs(float64(forwardSum - backwardSum)))

}

func lonelyInteger(a []int32) int32 {
	m := make(map[int32]int32)

	for _, value := range a {
		m[value]++
	}

	for key, value := range m {
		if value == 1 {
			return key
		}
	}

	return 0

}

func gradingStudents(grades []int32) []int32 {
	var result []int32
	for _, grade := range grades {
		if grade >= 38 {
			if grade%5 >= 3 {
				grade = grade + (5 - (grade % 5))
			}
		}
		result = append(result, grade)
	}
	return result

}

func matchingStrings(strings []string, queries []string) []int32 {
	m := make(map[string]int)
	var result []int32

	for _, value := range strings {
		m[value]++
	}

	for _, query := range queries {
		result = append(result, int32(m[query]))
	}
	return result

}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	counter := 0

	for i := 0; i < int(n); i++ {
		for j := i + 1; j < int(n); j++ {
			if (ar[i]+ar[j])%k == 0 {
				counter++
			}
		}
	}
	return int32(counter)

}

func isOneEditDifference(input1, input2 string) bool {
	if input1 == input2 {
		return true
	}
	runeArr2 := []rune(input2)
	var runeArr1 []rune

	for _, value := range input1 {
		runeArr1 = append(runeArr1, value)
	}

	// Check if removal of one character possible
	if checkInsert(runeArr1, runeArr2) {
		return true
	}

	// Check if insert of one character possible
	if checkInsert(runeArr2, runeArr1) {
		return true
	}

	// Check if replace of one character possible
	for i := 0; i < len(runeArr2); i++ {
		runeArrCopy1 := make([]rune, len(runeArr1))
		copy(runeArrCopy1, runeArr1)
		runeArrCopy2 := make([]rune, len(runeArr2))
		copy(runeArrCopy2, runeArr2)
		stringWithOneElementRemoved1 := string(delChar(runeArrCopy1, i))
		stringWithOneElementRemoved2 := string(delChar(runeArrCopy2, i))
		if stringWithOneElementRemoved1 == stringWithOneElementRemoved2 {
			return true
		}
	}

	return false
}

func checkInsert(runeArr1 []rune, runeArr2 []rune) bool {
	for i := 0; i < len(runeArr1); i++ {
		runeArrCopy := make([]rune, len(runeArr1))
		copy(runeArrCopy, runeArr1)
		stringWithOneElementRemoved := string(delChar(runeArrCopy, i))
		if stringWithOneElementRemoved == string(runeArr2) {
			return true
		}
	}
	return false
}

func isPermutationOfPalindrome(input string) bool {
	input = strings.ToUpper(input)
	input = strings.ReplaceAll(input, " ", "")

	var charSet int
	for _, value := range input {
		normalizedValue := value - 64
		if hasBit(charSet, uint(normalizedValue)) {
			charSet = clearBit(charSet, uint(normalizedValue))
		} else {
			charSet = setBit(charSet, uint(normalizedValue))
		}

	}
	mask := charSet - 1
	if charSet&mask == 0 {
		return true
	} else {
		return false
	}
}

func isPalindrome(input string) bool {
	input = strings.ToUpper(input)
	input = strings.ReplaceAll(input, " ", "")

	if len(input)%2 != 0 {
		input = string(delChar([]rune(input), len(input)/2))
	}

	if input[:len(input)/2] == reverse(input[len(input)/2:]) {
		return true
	} else {
		return false
	}
}

func hasUniqueCharacters(input string) bool {
	var charSet [128]bool
	for _, char := range input {
		if charSet[char] == true {
			return false
		} else {
			charSet[char] = true
		}
	}
	return true
}

func isPermutation(inputString1, inputString2 string) bool {
	inputString1 = strings.ToUpper(inputString1)
	inputString2 = strings.ToUpper(inputString2)
	charMap1 := getCharMapFromString(inputString1)
	charMap2 := getCharMapFromString(inputString2)

	if reflect.DeepEqual(charMap1, charMap2) {
		return true
	} else {
		return false
	}

}

func getCharMapFromString(inputString1 string) map[rune]int {
	m := make(map[rune]int)
	for _, char := range inputString1 {
		m[char] += 1
	}
	return m
}

func replaceSpacesInCharArray(input []rune) []rune {
	var output []rune
	for _, value := range input {
		if value == ' ' {
			output = append(output, '%')
			output = append(output, '2')
			output = append(output, '0')
		} else {
			output = append(output, value)
		}
	}
	return output
}

func compressString(input string) string {
	var compressedArr []rune
	lastChar := ' '
	var charCounter int64 = 0
	for _, value := range input {
		if lastChar != ' ' {
			if value != lastChar {
				compressedArr = addCompressedValue(compressedArr, lastChar, charCounter)
				lastChar = value
				charCounter = 1
			} else {
				charCounter++
			}
		} else {
			lastChar = value
			charCounter++
		}
	}
	compressedArr = addCompressedValue(compressedArr, lastChar, charCounter)

	return string(compressedArr)

}

func addCompressedValue(compressedArr []rune, lastChar int32, charCounter int64) []rune {
	compressedArr = append(compressedArr, lastChar)
	for _, value := range strconv.FormatInt(charCounter, 10) {
		compressedArr = append(compressedArr, value)
	}
	return compressedArr
}

func reverseMatrix(matrix [][]int) [][]int {
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	return matrix
}

func transposeMatrix(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func rotateMatrix(matrix [][]int) [][]int {
	return transposeMatrix(reverseMatrix(matrix))
}

func formatCamelCase(input []string) []string {
	var result []string
	for _, value := range input {
		var formattedWord string
		stringWithoutCommand := value[4:]
		if value[0] == 'S' {
			var finalArr []rune
			for i := 0; i < len(stringWithoutCommand); i++ {
				if unicode.IsUpper(rune(stringWithoutCommand[i])) {
					if i != 0 {
						finalArr = append(finalArr, ' ')
					}
					finalArr = append(finalArr, unicode.ToLower(rune(stringWithoutCommand[i])))
				} else {
					finalArr = append(finalArr, rune(stringWithoutCommand[i]))
				}
			}
			formattedWord = string(finalArr)
			formattedWord = strings.ReplaceAll(formattedWord, "(", "")
			formattedWord = strings.ReplaceAll(formattedWord, ")", "")
			result = append(result, formattedWord)

		} else {
			var finalArr []rune
			var spaceOccurred bool

			for _, char := range stringWithoutCommand {
				if char == ' ' {
					spaceOccurred = true
				} else {
					if spaceOccurred {
						finalArr = append(finalArr, unicode.ToUpper(char))
						spaceOccurred = false
					} else {
						finalArr = append(finalArr, char)
					}

				}

			}
			if value[2] == 'M' {
				finalArr = append(finalArr, '(')
				finalArr = append(finalArr, ')')
			} else if value[2] == 'C' {
				finalArr[0] = unicode.ToUpper(finalArr[0])
			}

			formattedWord = string(finalArr)
			result = append(result, formattedWord)
		}
	}
	return result
}
