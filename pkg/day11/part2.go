package day11

import (
	"fmt"
)

func SolvePart2(input string) string {
	monkeys := fromString(input)

	playRounds(10000, 1, &monkeys)

	return fmt.Sprint(businessFactor(monkeys))
}
