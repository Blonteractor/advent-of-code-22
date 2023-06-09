package day7

import (
	"fmt"
)

const (
	totalDiskSpace = 70000000
	updateSize     = 30000000
)

func SolvePart2(input string) string {
	root := parseInput(input)
	spaceNeededForUpdate := updateSize - (totalDiskSpace - root.CalculateSize())
	smallestToDelete := totalDiskSpace
	for _, dir := range root.Directories() {
		size := dir.CalculateSize()
		if size >= spaceNeededForUpdate && size < smallestToDelete {
			smallestToDelete = size
		}
	}

	return fmt.Sprint(smallestToDelete)
}
