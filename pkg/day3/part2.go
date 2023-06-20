package day3

import (
	"fmt"
	"strings"

	"github.com/Bonteractor/advent-of-code-22/pkg/util"
)

func SolvePart2(input string) string {
	ruckSacks := strings.Split(input, "\n")
	prioritySum := 0

	for i := 0; i < len(ruckSacks); i = i + 3 {
		first := util.SetFromString(ruckSacks[i])

		badge := intersection(intersection(first, util.SetFromString(ruckSacks[i+1])), util.SetFromString(ruckSacks[i+2]))

		for item := range badge {
			prioritySum += CalcPriority(item)
		}
	}

	return fmt.Sprint(prioritySum)
}
