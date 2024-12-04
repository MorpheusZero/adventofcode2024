package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func MainP2() {

	var leftList []int
	var rightList []int

	lines := strings.Split(PUZZLE_INPUT, "\n")
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		leftValue, _ := strconv.Atoi(parts[0])
		rightValue, _ := strconv.Atoi(parts[1])
		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	var RUNNING_TOTAL = 0

	for _, value := range leftList {

		// Get the number of occurances of the value in the left list
		occurances := getOccurancesOfIntFromArray(rightList, value)

		RUNNING_TOTAL += occurances * value
	}

	fmt.Println("Running total: ", RUNNING_TOTAL)
}

func getOccurancesOfIntFromArray(array []int, value int) int {
	count := 0
	for _, item := range array {
		if item == value {
			count++
		}
	}
	return count
}
