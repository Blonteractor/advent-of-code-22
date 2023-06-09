package day6

import (
	"fmt"
	"log"
)

func SolvePart2(input string) string {
	idx, err := findFirstUniqueSet(input, 14)
	if err != nil {
		log.Fatalf("%s", err)
	}

	return fmt.Sprint(idx)
}
