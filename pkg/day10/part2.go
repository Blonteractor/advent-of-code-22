package day10

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	crtWidth  = 40
	crtHeight = 6
	pixelLit  = "#"
	pixelDark = "."
)

func handleCylceDraw(rgx int, cycle int, screen [][]bool) {
	drawBufferPosition := (cycle - 1) % crtWidth
	drawBufferLine := (cycle - 1) / crtWidth
	if drawBufferPosition == rgx || drawBufferPosition == rgx+1 || drawBufferPosition == rgx-1 {
		screen[drawBufferLine][drawBufferPosition] = true
	}
}

func SolvePart2(input string) string {
	// rgx is the same as location of middle of sprite
	rgx := 1

	cycle := 1

	// Initialize the screen
	screen := make([][]bool, crtHeight)
	for i := range screen {
		screen[i] = make([]bool, crtWidth)
	}

	for _, v := range strings.Split(input, "\n") {
		instruction := strings.Split(v, " ")
		op := instruction[0]

		handleCylceDraw(rgx, cycle, screen)

		switch op {
		case "noop":
			cycle++
		case "addx":
			arg, err := strconv.Atoi(instruction[1])
			if err != nil {
				log.Fatalf("Trouble parsing input, arg for addx: %s", v)
			}
			cycle += 1
			handleCylceDraw(rgx, cycle, screen)
			cycle += 1
			rgx += arg
		default:
			log.Fatalf("Unkown operation: %s", op)
		}

	}

	var render string
	for i := range screen {
		for _, pixel := range screen[i] {
			if pixel {
				render += pixelLit
			} else {
				render += pixelDark
			}
		}
		render += "\n"
	}

	fmt.Println(render)
	return "SEE ABOVE RENDER"
}
