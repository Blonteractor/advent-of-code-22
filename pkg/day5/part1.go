package day5

import (
	"strings"
)

func SolvePart1(input string) string {
	s := strings.Split(input, "\n\n")
	stacks := parseCargoStacks(s[0])

	for _, instruction := range strings.Split(s[1], "\n") {
		instructions := craneInstruction(instruction)
		for i := 0; i < instructions.howMany; i++ {
			if stacks[instructions.from].Len() == 0 {
				break
			}
			cargo := stacks[instructions.from].PopBack()
			stacks[instructions.to].PushBack(cargo)
		}
	}

	return peekTop(stacks)
}
