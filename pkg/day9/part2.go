package day9

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	bodyLength = 10
	head       = 0
	tail       = bodyLength - 1
)

func SolvePart2(input string) string {
	visited := make(Set, 0)
	visited.Add(Position{})

	body := make([]Position, bodyLength)

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
			body[head].moveBy(direction_offset)
			for j := 0; j < bodyLength-1; j++ {
				body[j].moveTail(&body[j+1])
			}
			visited.Add(body[tail])
		}
	}

	return fmt.Sprint(len(visited))
}
