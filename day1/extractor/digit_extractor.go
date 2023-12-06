package extractor

import (
	"AoC23/internal/common/string_utils"
	"strconv"
	"unicode"
)

type DigitExtractor interface {
	ExtractFromString(input string) int
}

type FirstAndLastDigitExtractor struct{}

func (f FirstAndLastDigitExtractor) ExtractFromString(input string) int {
	return (getFirstDigitFromString(input) * 10) + getLastDigitFromString(input)
}

func getFirstDigitFromString(input string) int {
	for _, char := range input {
		if unicode.IsDigit(char) {
			value, _ := strconv.Atoi(string(char))
			return value
		}
	}

	return 0
}

func getLastDigitFromString(input string) int {
	for _, char := range string_utils.ReverseString(input) {
		if unicode.IsDigit(char) {
			value, _ := strconv.Atoi(string(char))
			return value
		}
	}

	return 0
}
