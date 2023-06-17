package day9

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func SolvePart1(input string) string {
	visited := make(Set, 0)
	visited.Add(Position{})
	var head Position
	var tail Position

	for _, v := range strings.Split(input, "\n") {
		instruction := strings.Split(v, " ")
		direction := instruction[0]
		magnitude, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatalf("Trouble parsing input for magnitude: %s", v)
		}

		var direction_offset Position
		switch direction {
		case "R":
			direction_offset.x = 1
		case "L":
			direction_offset.x = -1
		case "U":
			direction_offset.y = 1
		case "D":
			direction_offset.y = -1
		}

		for i := 0; i < magnitude; i++ {
			head.moveBy(direction_offset)
			head.moveTail(&tail)
			visited.Add(tail)
		}
	}

	return fmt.Sprint(len(visited))
}
