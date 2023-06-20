package day3

import (
	"fmt"
	"strings"

	"github.com/Bonteractor/advent-of-code-22/pkg/util"
)

func SolvePart1(input string) string {
	ruckSacks := strings.Split(input, "\n")
	prioritySum := 0

	for _, ruckSack := range ruckSacks {
		ruckSum := 0
		stackLen := len(ruckSack)

		firstCompartment := util.SetFromString(ruckSack[0 : stackLen/2])
		secondCompartment := util.SetFromString(ruckSack[stackLen/2 : stackLen])

		for item := range intersection(firstCompartment, secondCompartment) {
			ruckSum += CalcPriority(item)
		}
		prioritySum += ruckSum
	}

	return fmt.Sprint(prioritySum)
}
