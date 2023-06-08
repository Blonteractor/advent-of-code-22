package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Bonteractor/advent-of-code-22/pkg/day1"
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
	
	var result int
	
	// Call the appropriate solution function based on the day
	switch day {
	case 1:
		result = day1.SolvePart1(input)
		
	// Add cases for other days as needed
	default:
		log.Fatalf("Solution for day %d not implemented.", day)
	}
	
	fmt.Printf("Day %d Solution: %v\n", day, result)
}
