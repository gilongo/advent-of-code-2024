package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf("Day 3\n")

	var result int = 0
	data := string(utils.ReadInputFile("day3/input.txt"))
	reg := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	found := reg.MatchString(data)

	if found {
		foundData := reg.FindAllStringSubmatch(data, -1)
		regMul := regexp.MustCompile(`[0-9]{1,3}`)
		for _, mulStr := range foundData {
			factors := regMul.FindAllStringSubmatch(mulStr[0], -1)
			f1, err := strconv.Atoi(factors[0][0])
			utils.CheckErr(err)
			f2, err := strconv.Atoi(factors[1][0])
			utils.CheckErr(err)

			result += mul(f1, f2)
		}
	}

	fmt.Printf("Result: %d\n", result)
	Day3()
}

func mul(a, b int) int {
	return a * b
}
