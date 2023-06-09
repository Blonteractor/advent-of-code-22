package day3

import (
	"fmt"
	"strings"
)

func SolvePart1(input string) string {
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

	return fmt.Sprint(prioritySum)
}
