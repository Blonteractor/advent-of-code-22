package day1

import (
	"log"
	"strconv"
	"strings"
)

func SolvePart1(input string) int {
	elveList := strings.Split(input, "\n\n")
	max := 0
	
	for _, calorieList := range elveList {
		totalCalories := 0
		for _, calories := range strings.Split(calorieList, "\n") {
			c, err := strconv.Atoi(calories)
			if err != nil {
				log.Fatalf("Trouble parsing input: %s", calories)
			}
			totalCalories += c
		}
		if totalCalories > max {
			max = totalCalories
		}
	}
	
	return max;
}