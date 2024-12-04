package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MainP1() {

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

	var RUNNING_TOTAL = 0.0

	for len(leftList) > 0 || len(rightList) > 0 {

		// Get the smallest value from the left list
		smallestLeft := getSmallestIntFromArray(leftList)
		// Get the smallest value from the right list
		smallestRight := getSmallestIntFromArray(rightList)

		RUNNING_TOTAL += math.Abs(float64(smallestLeft - smallestRight))

		// Remove the smallest value from the left list
		leftList = removeIntFromArray(leftList, smallestLeft)

		// Remove the smallest value from the right list
		rightList = removeIntFromArray(rightList, smallestRight)

		fmt.Println("Items left to parse: ", len(leftList))
		// fmt.Println("Running total: ", RUNNING_TOTAL)
	}

	fmt.Println("Running total: ", RUNNING_TOTAL)
}
