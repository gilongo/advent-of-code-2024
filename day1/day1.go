package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
)

func main() {
	data := readInputFile("day1/input.txt")

	var firstList []int
	var secondList []int

	// split in two lists
	firstList, secondList = convertDataToSlices(data)

	// sort lists
	sort.Slice(firstList, func(i, j int) bool {
		return firstList[i] < firstList[j]
	})

	sort.Slice(secondList, func(i, j int) bool {
		return secondList[i] < secondList[j]
	})

	// calculate distance between points of two lists
	fmt.Println("Distance:", calculateDistanceSlices(firstList, secondList))

	// calculate similarity score
	fmt.Println("Similarity score:", calculateSimilarityScore(firstList, secondList))
}

func calculateSimilarityScore(firstList []int, secondList []int) int {
	similarityScore := 0
	alreadyUsed := map[int]int{}

	for i := 0; i < len(firstList); i++ {
		val, ok := alreadyUsed[firstList[i]]

		if ok {
			similarityScore += val * firstList[i]
			continue
		}

		ecountered := 0
		for j := 0; j < len(secondList); j++ {
			if firstList[i] == secondList[j] {
				ecountered++
			}
		}

		alreadyUsed[firstList[i]] = ecountered
		similarityScore += ecountered * firstList[i]
	}

	return similarityScore
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func calculateDistanceSlices(x []int, y []int) int {
	listDiff := 0
	for i := 0; i < len(x); i++ {
		listDiff += abs(x[i] - y[i])
	}

	return listDiff
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertDataToSlices(data []byte) ([]int, []int) {
	var firstList []int
	var secondList []int
	var breakLine string

	if runtime.GOOS == "windows" {
		breakLine = "\r\n"
	} else {
		breakLine = "\n"
	}

	parts := bytes.Split(data, []byte(breakLine))
	for _, part := range parts {
		part1, part2, _ := bytes.Cut(part, []byte("   "))

		intPart1, err := strconv.Atoi(string(part1))
		check(err)

		intPart2, err := strconv.Atoi(string(part2))
		check(err)

		firstList = append(firstList, intPart1)
		secondList = append(secondList, intPart2)
	}

	return firstList, secondList
}

func readInputFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}
