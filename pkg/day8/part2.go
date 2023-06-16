package day8

import "fmt"

func SolvePart2(input string) string {
	forestGrid := parseGrid(input)
	highestScenicScore := 0

	for i := 0; i < len(forestGrid); i++ {
		for j := 0; j < len(forestGrid[i]); j++ {
			scenicScore := visionFromUp(i, j, forestGrid) * visionFromDown(i, j, forestGrid) * visionFromRight(i, j, forestGrid) * visionFromLeft(i, j, forestGrid)
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return fmt.Sprint(highestScenicScore)
}

func visionFromLeft(x int, y int, forestGrid Grid) int {
	vision := 0
	for i := y - 1; i >= 0; i-- {
		vision++
		if forestGrid[x][i] >= forestGrid[x][y] {
			return vision
		}
	}
	return vision
}

func visionFromRight(x int, y int, forestGrid Grid) int {
	vision := 0
	for i := y + 1; i < len(forestGrid[x]); i++ {
		vision++
		if forestGrid[x][i] >= forestGrid[x][y] {
			return vision
		}
	}
	return vision
}

func visionFromUp(x int, y int, forestGrid Grid) int {
	vision := 0
	for i := x - 1; i >= 0; i-- {
		vision++
		if forestGrid[i][y] >= forestGrid[x][y] {
			return vision
		}
	}
	return vision
}

func visionFromDown(x int, y int, forestGrid Grid) int {
	vision := 0
	for i := x + 1; i < len(forestGrid); i++ {
		vision++
		if forestGrid[i][y] >= forestGrid[x][y] {
			return vision
		}
	}
	return vision
}
