package day4

import (
	"fmt"
	"strings"
)

func SolvePart2(input string) string {
	result := 0
	for _, v := range strings.Split(input, "\n") {
		assignementPair := strings.Split(v, ",")

		range1 := numberRangeFromString(assignementPair[0])
		range2 := numberRangeFromString(assignementPair[1])

		if range1.isOverlapping(range2) || range2.isOverlapping(range1) {
			result++
		}
	}

	return fmt.Sprint(result)
}
