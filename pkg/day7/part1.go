package day7

import (
	"fmt"
)

const maxSizeToDelete = 100000

func SolvePart1(input string) string {
	root := parseInput(input)
	sumOfDeletableDir := 0

	for _, dir := range root.Directories() {
		size := dir.CalculateSize()
		if size <= maxSizeToDelete {
			sumOfDeletableDir += size
		}
	}

	return fmt.Sprint(sumOfDeletableDir)
}
