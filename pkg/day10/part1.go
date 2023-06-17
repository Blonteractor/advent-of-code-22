package day10

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	baseCycle = 20
	lastCycle = 220
	incriment = 40
)

func handleCylceSignal(rgx int, cycle int, i *int, signal_strengths *[]int) {
	signal_strength := cycle * rgx
	if cycle == baseCycle+*i*incriment {
		*signal_strengths = append(*signal_strengths, signal_strength)
		*i++
	}
}

func SolvePart1(input string) string {
	rgx := 1
	cycle := 1
	signal_strengths := make([]int, 0)
	i := 0

	for _, v := range strings.Split(input, "\n") {
		instruction := strings.Split(v, " ")
		op := instruction[0]

		handleCylceSignal(rgx, cycle, &i, &signal_strengths)

		switch op {
		case "noop":
			cycle++
		case "addx":
			arg, err := strconv.Atoi(instruction[1])
			if err != nil {
				log.Fatalf("Trouble parsing input, arg for addx: %s", v)
			}
			cycle++
			handleCylceSignal(rgx, cycle, &i, &signal_strengths)
			cycle++
			rgx += arg
		default:
			log.Fatalf("Unkown operation: %s", op)
		}

	}

	sum := 0
	for _, v := range signal_strengths {
		sum += v
	}

	return fmt.Sprint(sum)
}
