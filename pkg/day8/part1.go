package day8

import "fmt"

func SolvePart1(input string) string {
	forestGrid := parseGrid(input)

	// the trees at the edge are already visible
	visibleCount := 2*(len(forestGrid)+len(forestGrid[0])) - 4
	for i := 1; i < len(forestGrid)-1; i++ {
		for j := 1; j < len(forestGrid[i])-1; j++ {
			if isVisibleFromUp(i, j, forestGrid) || isVisibleFromDown(i, j, forestGrid) || isVisibleFromRight(i, j, forestGrid) || isVisibleFromLeft(i, j, forestGrid) {
				visibleCount++
			}
		}
	}

	return fmt.Sprint(visibleCount)
}

func isVisibleFromLeft(x int, y int, forestGrid Grid) bool {
	for i := y - 1; i >= 0; i-- {
		if forestGrid[x][i] >= forestGrid[x][y] {
			return false
		}
	}
	return true
}

func isVisibleFromRight(x int, y int, forestGrid Grid) bool {
	for i := y + 1; i < len(forestGrid[x]); i++ {
		if forestGrid[x][i] >= forestGrid[x][y] {
			return false
		}
	}
	return true
}

func isVisibleFromUp(x int, y int, forestGrid Grid) bool {
	for i := x - 1; i >= 0; i-- {
		if forestGrid[i][y] >= forestGrid[x][y] {
			return false
		}
	}
	return true
}

func isVisibleFromDown(x int, y int, forestGrid Grid) bool {
	for i := x + 1; i < len(forestGrid); i++ {
		if forestGrid[i][y] >= forestGrid[x][y] {
			return false
		}
	}
	return true
}
