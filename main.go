package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Bonteractor/advent-of-code-22/pkg/day1"
	"github.com/Bonteractor/advent-of-code-22/pkg/day2"
	"github.com/Bonteractor/advent-of-code-22/pkg/day3"
	"github.com/Bonteractor/advent-of-code-22/pkg/day4"
	"github.com/Bonteractor/advent-of-code-22/pkg/day5"
	"github.com/Bonteractor/advent-of-code-22/pkg/day6"
	"github.com/Bonteractor/advent-of-code-22/pkg/day7"
	"github.com/Bonteractor/advent-of-code-22/pkg/day8"
)

func main() {
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

	var solution Solution

	// Call the appropriate solution function based on the day
	switch day {
	case 1:
		solution = Solution{
			Part1: day1.SolvePart1(string(input)),
			Part2: day1.SolvePart2(string(input)),
		}
	case 2:
		solution = Solution{
			Part1: day2.SolvePart1(string(input)),
			Part2: day2.SolvePart2(string(input)),
		}
	case 3:
		solution = Solution{
			Part1: day3.SolvePart1(string(input)),
			Part2: day3.SolvePart2(string(input)),
		}

	case 4:
		solution = Solution{
			Part1: day4.SolvePart1(string(input)),
			Part2: day4.SolvePart2(string(input)),
		}

	case 5:
		solution = Solution{
			Part1: day5.SolvePart1(string(input)),
			Part2: day5.SolvePart2(string(input)),
		}

	case 6:
		solution = Solution{
			Part1: day6.SolvePart1(string(input)),
			Part2: day6.SolvePart2(string(input)),
		}

	case 7:
		solution = Solution{
			Part1: day7.SolvePart1(string(input)),
			Part2: day7.SolvePart2(string(input)),
		}

	case 8:
		solution = Solution{
			Part1: day8.SolvePart1(string(input)),
			Part2: day8.SolvePart2(string(input)),
		}

	// Add cases for other days as needed
	default:
		log.Fatalf("Solution for day %d not implemented.", day)
	}

	fmt.Printf("Day %d Solution: %v\n", day, solution)
}

type Solution struct {
	Part1 string
	Part2 string
}
