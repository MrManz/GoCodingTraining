package GoCodingTraining

import (
	"reflect"
	"strings"
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
