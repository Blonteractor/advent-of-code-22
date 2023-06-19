package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Bonteractor/advent-of-code-22/pkg/day1"
	"github.com/Bonteractor/advent-of-code-22/pkg/day10"
	"github.com/Bonteractor/advent-of-code-22/pkg/day11"
	"github.com/Bonteractor/advent-of-code-22/pkg/day12"
	"github.com/Bonteractor/advent-of-code-22/pkg/day2"
	"github.com/Bonteractor/advent-of-code-22/pkg/day3"
	"github.com/Bonteractor/advent-of-code-22/pkg/day4"
	"github.com/Bonteractor/advent-of-code-22/pkg/day5"
	"github.com/Bonteractor/advent-of-code-22/pkg/day6"
	"github.com/Bonteractor/advent-of-code-22/pkg/day7"
	"github.com/Bonteractor/advent-of-code-22/pkg/day8"
	"github.com/Bonteractor/advent-of-code-22/pkg/day9"
)

type PartSolver func(string) string
type Solver struct {
	part1 PartSolver
	part2 PartSolver
}

func (s Solver) getSolution(input string) Solution {
	return Solution{
		Part1: s.part1(input),
		Part2: s.part2(input),
	}
}

func main() {
	solutions := []Solver{
		{
			part1: day1.SolvePart1,
			part2: day1.SolvePart2,
		},
		{
			part1: day2.SolvePart1,
			part2: day2.SolvePart2,
		},
		{
			part1: day3.SolvePart1,
			part2: day3.SolvePart2,
		},
		{
			part1: day4.SolvePart1,
			part2: day4.SolvePart2,
		},
		{
			part1: day5.SolvePart1,
			part2: day5.SolvePart2,
		},
		{
			part1: day6.SolvePart1,
			part2: day6.SolvePart2,
		},
		{
			part1: day7.SolvePart1,
			part2: day7.SolvePart2,
		},
		{
			part1: day8.SolvePart1,
			part2: day8.SolvePart2,
		},
		{
			part1: day9.SolvePart1,
			part2: day9.SolvePart2,
		},
		{
			part1: day10.SolvePart1,
			part2: day10.SolvePart2,
		},
		{
			part1: day11.SolvePart1,
			part2: day11.SolvePart2,
		},
		{
			part1: day12.SolvePart1,
			part2: day12.SolvePart2,
		},
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Ensure a command-line argument is provided for the day
	if len(os.Args) < 2 {
		log.Fatal("Please provide the day as a command-line argument.")
	}

	// Parse the day argument
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid day argument: %s", os.Args[1])
	}

	// Get the input file path for the specified day
	inputFilePath := filepath.Join(dir, "inputs", fmt.Sprintf("day%d.txt", day))

	// Check if the input file exists
	_, err = os.Stat(inputFilePath)
	if os.IsNotExist(err) {
		log.Fatalf("Input file for day %d not found.", day)
	}

	// Read the input file
	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if day > len(solutions) {
		log.Fatalf("Solution for day %d not implemented.", day)
	}

	solution := solutions[day-1].getSolution(string(input))

	fmt.Printf("Day %d Solution: %v\n", day, solution)
}

type Solution struct {
	Part1 string
	Part2 string
}
