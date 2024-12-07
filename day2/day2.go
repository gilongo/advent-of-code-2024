package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var reportLists [][]int

	data := utils.ReadInputFile("day2/input.txt")
	safeReports := 0

	reportLists = convertDataToReportLists(data)

	for _, report := range reportLists {
		var reportSerieBehaviour []bool
		safe := true

		for i := range report {
			if i == 0 {
				continue
			}

			reportSerieBehaviour = append(reportSerieBehaviour, report[i] > report[i-1])
			if i-1 > 0 {
				if reportSerieBehaviour[i-1] != reportSerieBehaviour[i-2] {
					safe = false
					break
				}
			}

			reportLevel := utils.Abs(report[i] - report[i-1])
			if reportLevel > 3 || reportLevel == 0 {
				safe = false
				break
			}
		}

		if safe {
			safeReports++
		}
	}

	fmt.Println("Safe reports:", safeReports)
	fmt.Println("More Safer reports:", Part2())
}

func convertDataToReportLists(data []byte) [][]int {
	var reports [][]int
	for _, line := range utils.SplitData(data, utils.GetBreakLineToken()) {

		var reportLine []int
		for _, val := range strings.Split(string(line), " ") {
			reportVal, err := strconv.Atoi(val)
			utils.CheckErr(err)

			reportLine = append(reportLine, reportVal)
		}

		reports = append(reports, reportLine)
	}

	return reports
}
