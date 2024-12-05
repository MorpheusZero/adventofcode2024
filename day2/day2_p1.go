package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	ReportTypeASC     = "ASC"
	ReportTypeDESC    = "DESC"
	ReportTypeUNKNOWN = "UNKNOWN"
)

type Report struct {
	ID          int
	Values      []int
	IsSafe      bool
	ReportType  string
	LevelIssues int
	Spicy       bool
	Hultipoo    bool
}

func MainP1() {
	var SAFE_REPORTS_COUNT = 0

	lines := strings.Split(PUZZLE_INPUT, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		var report = Report{}
		report.Values = make([]int, 0)
		report.IsSafe = true
		report.ReportType = ReportTypeUNKNOWN
		report.LevelIssues = 0
		var value = 0
		for _, part := range parts {
			if part != "" {
				value, _ = strconv.Atoi(part)
				report.Values = append(report.Values, value)
			}
		}
		var isSafe = CheckIfReportValuesAreSafe(&report)
		if isSafe {
			SAFE_REPORTS_COUNT++
		}
	}

	fmt.Print("Safe reports count: " + strconv.Itoa(SAFE_REPORTS_COUNT))

}

func CheckIfReportValuesAreSafe(report *Report) bool {
	for i, value := range report.Values {
		// On the first iteration, skip the check
		if i == 0 {
			continue
		}

		// On the second iteration, let's try to determine the report type whether ASC or DESC
		if i == 1 {
			if value < report.Values[i-1] {
				report.ReportType = ReportTypeDESC
			} else {
				report.ReportType = ReportTypeASC
			}
		}

		// If the report type is DESC, then the value should be less than the previous value
		if report.ReportType == ReportTypeDESC {
			if value < report.Values[i-1] {
				var diff = math.Abs(float64(report.Values[i-1] - value))
				if diff == 0 || diff > 3 {
					report.IsSafe = false
				}
			} else {
				report.IsSafe = false
			}
		}

		// If the report type is ASC, then the value should be greater than the previous value
		if report.ReportType == ReportTypeASC {
			if value > report.Values[i-1] {
				var diff = math.Abs(float64(report.Values[i-1] - value))
				if diff == 0 || diff > 3 {
					report.IsSafe = false
				}
			} else {
				report.IsSafe = false
			}
		}
	}
	return report.IsSafe
}
