package day5

import (
	"strings"

	"github.com/gammazero/deque"
)

func SolvePart2(input string) string {
	s := strings.Split(input, "\n\n")
	stacks := parseCargoStacks(s[0])

	for _, instruction := range strings.Split(s[1], "\n") {
		instructions := craneInstruction(instruction)
		var payload deque.Deque[string]

		for i := 0; i < instructions.howMany; i++ {
			if stacks[instructions.from].Len() == 0 {
				break
			}
			cargo := stacks[instructions.from].PopBack()
			payload.PushBack(cargo)
		}
		for payload.Len() != 0 {
			stacks[instructions.to].PushBack(payload.PopBack())
		}
	}

	return peekTop(stacks)
}
