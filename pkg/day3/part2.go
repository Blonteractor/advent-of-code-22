package day3

import (
	"fmt"
	"strings"
)

func SolvePart2(input string) string {
	ruckSacks := strings.Split(input, "\n")
	prioritySum := 0

	for i := 0; i < len(ruckSacks); i = i + 3 {
		first := SetFromString(ruckSacks[i])

		intersection := first.Intersection(SetFromString(ruckSacks[i+1]))
		badge := intersection.Intersection(SetFromString(ruckSacks[i+2]))

		for item := range badge {
			prioritySum += CalcPriority(item)
		}
	}

	return fmt.Sprint(prioritySum)
}
