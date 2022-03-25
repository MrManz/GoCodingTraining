package GoCodingTraining

import (
	"reflect"
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
	return (val > 0)
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
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
			if value[3] == 'M' {

			}
		}
	}
	return result
}

func formatCamelCase(input string) string {
	return ""
}
