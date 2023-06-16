package day8

import (
	"log"
	"strconv"
	"strings"
)

type Grid [][]int

func parseGrid(input string) Grid {
	inputGrid := strings.Split(input, "\n")
	parsedGrid := make([][]int, len(inputGrid))

	for i, line := range inputGrid {
		heights := strings.Split(line, "")
		row := make([]int, len(heights))
		for j, treeHeight := range heights {
			height, err := strconv.Atoi(treeHeight)
			if err != nil {
				log.Fatalf("Trouble parsing input at tree height: %s", treeHeight)
			}
			row[j] = height
		}
		parsedGrid[i] = row
	}

	return parsedGrid
}
