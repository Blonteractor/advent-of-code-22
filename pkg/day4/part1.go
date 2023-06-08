package day4

import (
	"strings"
)

func SolvePart1(input string) int {
	result := 0
	for _, v := range strings.Split(input, "\n") {
		assignementPair := strings.Split(v, ",")

		range1 := numberRangeFromString(assignementPair[0])
		range2 := numberRangeFromString(assignementPair[1])

		if range1.isContainedIn(range2) || range2.isContainedIn(range1) {
			result++
		}
	}

	return result
}
