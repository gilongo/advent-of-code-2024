package main

import (
	"advent-of-code-2024/utils"
)

func Part2() int {
	data := utils.ReadInputFile("day2/input.txt")
	reportLists := convertDataToReportLists(data)
	safe := 0

	for _, reportList := range reportLists {
		if isSafe(reportList) {
			safe++
		}
	}

	return safe
}

func isStrictlyIncreasing(sequence []int) bool {
	for i := 1; i < len(sequence); i++ {
		if sequence[i] <= sequence[i-1] || (sequence[i]-sequence[i-1] > 3) {
			return false
		}
	}
	return true
}

func isStrictlyDecreasing(sequence []int) bool {
	for i := 1; i < len(sequence); i++ {
		if sequence[i] >= sequence[i-1] || (sequence[i-1]-sequence[i] > 3) {
			return false
		}
	}
	return true
}

func isSafe(sequence []int) bool {
	if isStrictlyIncreasing(sequence) || isStrictlyDecreasing(sequence) {
		return true
	}

	for i := 0; i < len(sequence); i++ {
		newSequence := make([]int, len(sequence))
		copy(newSequence, sequence)
		newSequence = append(newSequence[:i], newSequence[i+1:]...)

		if isStrictlyIncreasing(newSequence) || isStrictlyDecreasing(newSequence) {
			return true
		}
	}

	return false
}
