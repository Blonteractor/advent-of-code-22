package day3

import (
	"strings"
)

func SolvePart1(input string) int {
	ruckSacks := strings.Split(input, "\n")
	prioritySum := 0

	for _, ruckSack := range ruckSacks {
		ruckSum := 0
		stackLen := len(ruckSack)

		firstCompartment := SetFromString(ruckSack[0 : stackLen/2])
		secondCompartment := SetFromString(ruckSack[stackLen/2 : stackLen])

		for item := range firstCompartment.Intersection(secondCompartment) {
			ruckSum += CalcPriority(item)
		}
		prioritySum += ruckSum
	}

	return prioritySum
}
