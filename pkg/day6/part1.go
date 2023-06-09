package day6

import (
	"fmt"
	"log"
)

func SolvePart1(input string) string {
	idx, err := findFirstUniqueSet(input, 4)
	if err != nil {
		log.Fatalf("%s", err)
	}

	return fmt.Sprint(idx)
}
