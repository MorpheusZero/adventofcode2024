package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var allReportsWithRemovedIndex []Report = []Report{}

func MainP2() {
	var SAFE_REPORTS_COUNT = 0

	lines := strings.Split(PUZZLE_INPUT, "\n")
	for idx, line := range lines {
		parts := strings.Split(line, " ")
		var report = Report{}
		report.ID = idx
		report.Values = make([]int, 0)
		report.IsSafe = true
		report.ReportType = ReportTypeUNKNOWN
		report.LevelIssues = 0
		report.Spicy = false
		var value = 0
		for _, part := range parts {
			if part != "" {
				value, _ = strconv.Atoi(part)
				report.Values = append(report.Values, value)
			}
		}
		var isSafe = CheckIfReportValuesAreSafeP2(report)
		if isSafe {
			SAFE_REPORTS_COUNT++
		}
	}

	for _, report := range allReportsWithRemovedIndex {
		var isSafe = CheckIfReportValuesAreSafeP2(report)
		if isSafe {
			SAFE_REPORTS_COUNT++
		}
	}

	fmt.Print("Safe reports count: " + strconv.Itoa(SAFE_REPORTS_COUNT))

}

func CheckIfReportValuesAreSafeP2(report Report) bool {

	for i, value := range report.Values {
		// On the first iteration, skip the check
		if i == 0 {
			continue
		}

		if report.ID == 69 {
			fmt.Println(report)
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
					report.LevelIssues++
					report.IsSafe = false
					if !report.Hultipoo {
						var newArrayValues = removeIndexFromArray(report.Values, i-1)
						allReportsWithRemovedIndex = append(allReportsWithRemovedIndex, Report{ID: report.ID, Values: newArrayValues, IsSafe: true, ReportType: report.ReportType, Spicy: true})
						report.Hultipoo = true
					}
				}
			} else {
				report.LevelIssues++
				report.IsSafe = false
				if !report.Hultipoo {
					var newArrayValues = removeIndexFromArray(report.Values, i-1)
					allReportsWithRemovedIndex = append(allReportsWithRemovedIndex, Report{ID: report.ID, Values: newArrayValues, IsSafe: true, ReportType: report.ReportType, Spicy: true})
					report.Hultipoo = true
				}
			}
		}

		// If the report type is ASC, then the value should be greater than the previous value
		if report.ReportType == ReportTypeASC {
			if value > report.Values[i-1] {
				var diff = math.Abs(float64(report.Values[i-1] - value))
				if diff == 0 || diff > 3 {
					report.LevelIssues++
					report.IsSafe = false
					if !report.Hultipoo {
						var newArrayValues = removeIndexFromArray(report.Values, i-1)
						if report.ID == 69 {
							fmt.Println(newArrayValues)
						}
						allReportsWithRemovedIndex = append(allReportsWithRemovedIndex, Report{ID: report.ID, Values: newArrayValues, IsSafe: true, ReportType: report.ReportType, Spicy: true})
						report.Hultipoo = true
					}
				}
			} else {
				report.LevelIssues++
				report.IsSafe = false
				if !report.Hultipoo {
					var newArrayValues = removeIndexFromArray(report.Values, i-1)
					if report.ID == 69 {
						fmt.Println(newArrayValues)
					}
					allReportsWithRemovedIndex = append(allReportsWithRemovedIndex, Report{ID: report.ID, Values: newArrayValues, IsSafe: true, ReportType: report.ReportType, Spicy: true})
					report.Hultipoo = true
				}
			}
		}
	}

	// if report.ID == 69 {
	// 	fmt.Println(report)
	// }

	// if report.Spicy && !report.IsSafe {
	// 	fmt.Println(report)
	// }

	return report.IsSafe
}

func removeIndexFromArray(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}
