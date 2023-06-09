package day5

import (
	"log"
	"strconv"
	"strings"

	"github.com/gammazero/deque"
)

type CraneInstruction struct {
	howMany int
	to      int
	from    int
}

func parseCargoStacks(input string) []deque.Deque[string] {
	stackConfiguration := strings.Split(input, "\n")
	stackNumbering := strings.Split(stackConfiguration[len(stackConfiguration)-1], "   ")

	numberOfStacks, err := strconv.Atoi(string(stackNumbering[len(stackNumbering)-1][0]))
	if err != nil {
		log.Fatalf("Trouble parsing input, number of stacks: %s", stackNumbering)
	}

	stacks := make([]deque.Deque[string], numberOfStacks)
	for _, boxLine := range stackConfiguration[:len(stackConfiguration)-1] {
		for i := 0; i < numberOfStacks; i++ {
			cargo := boxLine[4*i+1]
			if cargo != ' ' {
				stacks[i].PushFront(string(cargo))
			}
		}
	}

	return stacks
}

func craneInstruction(input string) CraneInstruction {
	s := strings.Split(input, " ")
	howMany, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatalf("Trouble parsing input, how many: %s", input)
	}

	from, err := strconv.Atoi(s[3])
	if err != nil {
		log.Fatalf("Trouble parsing input, from: %s", input)
	}
	from--

	to, err := strconv.Atoi(s[5])
	if err != nil {
		log.Fatalf("Trouble parsing input, to: %s", input)
	}
	to--

	return CraneInstruction{
		howMany: howMany,
		to:      to,
		from:    from,
	}
}

func peekTop(stacks []deque.Deque[string]) string {
	var result string
	for _, v := range stacks {
		top := v.PopBack()
		result += top
	}
	return result
}
