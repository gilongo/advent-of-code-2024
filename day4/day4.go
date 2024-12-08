package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"regexp"
)

func main() {
	searchXmasWord()
}

func searchXmasWord() {
	data := utils.ReadInputFile("day4/input.txt")

	horizontal := horizonalSearch(data)
	vertical := verticalSearch(data)
	diagonal := diagonalSearch(data)

	fmt.Printf("horizontal: %d, vertical: %d, diagonal: %d\n", horizontal, vertical, diagonal)
	fmt.Println("Total XMAS|SAMX words:", horizontal+vertical+diagonal)
}

func horizonalSearch(data []byte) int {
	horizontalBytes := [][]byte{}
	horizontalBytes = utils.SplitData(data, utils.GetBreakLineToken())

	horizontalStrings := []string{}
	for _, bytes := range horizontalBytes {
		horizontalStrings = append(horizontalStrings, string(bytes))
	}

	numXmasWords := searchXmasWords(horizontalStrings)

	return numXmasWords
}

func verticalSearch(data []byte) int {
	horizontalBytes := [][]byte{}
	horizontalBytes = utils.SplitData(data, utils.GetBreakLineToken())

	horizontalStrings := []string{}
	for _, bytes := range horizontalBytes {
		horizontalStrings = append(horizontalStrings, string(bytes))
	}

	numberOfColumns := len(horizontalStrings[0])
	numberOfRows := len(horizontalStrings)

	verticalStrings := []string{}
	verticalStrings = make([]string, numberOfColumns)

	// building vertical slices
	for c := 0; c < numberOfColumns; c++ {
		for r := 0; r < numberOfRows; r++ {
			verticalStrings[c] += string(horizontalStrings[r][c])
		}
	}

	numXmasWords := searchXmasWords(verticalStrings)

	return numXmasWords
}

func diagonalSearch(data []byte) int {
	horizontalBytes := [][]byte{}
	horizontalBytes = utils.SplitData(data, utils.GetBreakLineToken())

	horizontalStrings := []string{}
	for _, bytes := range horizontalBytes {
		horizontalStrings = append(horizontalStrings, string(bytes))
	}

	numberOfColumns := len(horizontalStrings[0])
	numberOfRows := len(horizontalStrings)

	// build right shifted strings
	numberOfShiftedColumns := numberOfColumns + numberOfRows - 1
	stringLen := len(horizontalStrings[0])
	rightShiftedStrings := []string{}
	rightShiftedStrings = make([]string, numberOfShiftedColumns)

	for r := 0; r < numberOfRows; r++ {
		prependString := generateFillString(r)
		appendString := generateFillString((numberOfShiftedColumns - stringLen) - r)

		if prependString != "" {
			rightShiftedStrings[r] += prependString
		}

		rightShiftedStrings[r] += horizontalStrings[r]

		if appendString != "" {
			rightShiftedStrings[r] += appendString
		}
	}

	// build left shifted strings
	leftShiftedStrings := []string{}
	leftShiftedStrings = make([]string, numberOfShiftedColumns)

	for r := 0; r < numberOfRows; r++ {
		prependString := generateFillString((numberOfShiftedColumns - stringLen) - r)
		appendString := generateFillString(r)

		if prependString != "" {
			leftShiftedStrings[r] += prependString
		}

		leftShiftedStrings[r] += horizontalStrings[r]

		if appendString != "" {
			leftShiftedStrings[r] += appendString
		}
	}

	// build vertical strings
	verticalStrings := []string{}
	verticalStrings = make([]string, numberOfShiftedColumns*2)

	i := 0
	// building vertical slices from rightShiftedStrings
	for c := 0; c < len(rightShiftedStrings); c++ {
		for r := 0; r < numberOfRows; r++ {
			verticalStrings[i] += string(rightShiftedStrings[r][c])
		}
		i++
	}

	// building vertical slices from leftShiftedStrings
	for c := 0; c < len(leftShiftedStrings); c++ {
		for r := 0; r < numberOfRows; r++ {
			verticalStrings[i] += string(leftShiftedStrings[r][c])
		}
		i++
	}

	numXmasWords := searchXmasWords(verticalStrings)

	return numXmasWords
}

func generateFillString(characters int) string {
	fillString := ""
	if characters > 0 {
		for i := 0; i < characters; i++ {
			fillString += "."
		}
	}
	return fillString
}

func searchXmasWords(input []string) int {

	xmasRegExp := regexp.MustCompile("XMAS")
	samxRegExp := regexp.MustCompile(`SAMX`)

	var foundXmasStrings []string
	for _, strings := range input {
		if xmasRegExp.FindString(strings) != "" {
			for _, match := range xmasRegExp.FindAllString(strings, -1) {
				foundXmasStrings = append(foundXmasStrings, match)
			}
		}
		if samxRegExp.FindString(strings) != "" {
			for _, match := range samxRegExp.FindAllString(strings, -1) {
				foundXmasStrings = append(foundXmasStrings, match)
			}
		}
	}

	return len(foundXmasStrings)
}
