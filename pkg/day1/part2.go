package day1

import (
	"log"
	"strconv"
	"strings"
)

func SolvePart2(input string) int {
	elveList := strings.Split(input, "\n\n")
	first, second, third := 0, 0, 0
	
	for _, calorieList := range elveList {
		totalCalories := 0
		for _, calories := range strings.Split(calorieList, "\n") {
			c, err := strconv.Atoi(calories)
			if err != nil {
				log.Fatalf("Trouble parsing input: %s", calories)
			}
			totalCalories += c
		}
		
		// Push every rank up according to case
		if totalCalories > first {
			third = second
			second = first
			first = totalCalories
		} else if (totalCalories > second && totalCalories != first) {
			third = second 
			second = totalCalories
		} else if (totalCalories > third && totalCalories != second) {
			third = totalCalories
		}
	}
	
	return first + second + third;
}