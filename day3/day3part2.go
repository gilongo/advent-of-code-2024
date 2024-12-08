package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Day3() {
	data := utils.ReadInputFile("day3/input.txt")
	reg := regexp.MustCompile(`((?:don't|do|mul)\(\d*,?\d*\))`)
	//factors := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	allStrings := reg.FindAllString(string(data), -1)
	fmt.Println(allStrings)

	enable := true
	var enabledStrings = make([]string, 0, len(allStrings))
	for _, str := range allStrings {
		if str == "do()" {
			enable = true
		} else if str == "don't()" {
			enable = false
		} else {
			if enable {
				enabledStrings = append(enabledStrings, str)
			}
		}
	}

	result := 0
	regMul := regexp.MustCompile(`[0-9]{1,3}`)
	for _, mulStr := range enabledStrings {
		factors := regMul.FindAllStringSubmatch(mulStr, -1)
		f1, err := strconv.Atoi(factors[0][0])
		utils.CheckErr(err)
		f2, err := strconv.Atoi(factors[1][0])
		utils.CheckErr(err)

		result += mul(f1, f2)
	}

	fmt.Println(result)
}
