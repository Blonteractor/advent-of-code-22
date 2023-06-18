package day11

import (
	"fmt"
)

func SolvePart1(input string) string {
	monkeys := fromString(input)

	playRounds(20, 3, &monkeys)

	return fmt.Sprint(businessFactor(monkeys))
}
