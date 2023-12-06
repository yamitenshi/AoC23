package main

import (
	"AoC23/day1/extractor"
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Specify an input file!")
	}

	readFile, err := os.Open(args[1])

	if err != nil {
		panic("File could not be opened!")
	}

	fileScanner := bufio.NewScanner(readFile)
	digitExtractor := extractor.FirstAndLastDigitExtractor{}
	result := 0

	for fileScanner.Scan() {
		inputLine := fileScanner.Text()
		result += digitExtractor.ExtractFromString(inputLine)
	}

	fmt.Printf("Result: %d\n", result)
}
