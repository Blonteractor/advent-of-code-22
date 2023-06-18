package day11

import (
	"log"
	"strconv"
	"strings"

	"github.com/gammazero/deque"
)

type WorryLevel uint64

const (
	itemListStart = 18
	operatorIdx   = 23
	operandLoc    = 7
	divisorLoc    = 5
	ifTrueLoc     = 9
	ifFalseLoc    = ifTrueLoc
)

type Monkey struct {
	items          deque.Deque[uint64]
	operation      func(uint64) uint64
	divisor        int
	ifTrue         *Monkey
	ifFalse        *Monkey
	itemsInspected int
}

func (m Monkey) test(worry uint64) bool {
	return worry%uint64(m.divisor) == 0
}

func (m *Monkey) inspectItems(worryFactor int, optimizationFactor int) {
	for m.items.Len() != 0 {
		// Seoy theorem
		// (ab + c) % k = [(a % j*k)(b % j*k) + c % j*k]  % k
		// where j is any integer
		// so what we are gonna do is take % of some factor, this will preserve the behavior when we take
		// mod by any other number while keeping worry lvl small
		// we can do this becuase we dont care about the value of the worry lvl, just the result of the test
		// we choose the lcm of the divisors as the factor
		new_worry := (m.operation(m.items.PopFront()) / uint64(worryFactor)) % uint64(optimizationFactor)
		if m.test(new_worry) {
			m.ifTrue.items.PushBack(new_worry)
		} else {
			m.ifFalse.items.PushBack(new_worry)
		}
		m.itemsInspected++
	}
}

func fromString(input string) []Monkey {
	inputs := strings.Split(input, "\n\n")
	monkeys := make([]Monkey, len(inputs))

	for i, v := range inputs {
		monkeyInfo := strings.Split(v, "\n")
		currentMonkey := &monkeys[i]

		for _, v := range strings.Split(monkeyInfo[1][itemListStart:], ", ") {
			currentMonkey.items.PushBack(parseToInt(v))
		}

		operandStr := strings.Split(monkeyInfo[2], " ")[operandLoc]
		operator := monkeyInfo[2][operatorIdx]
		currentMonkey.ifTrue = &monkeys[parseToInt(strings.Split(monkeyInfo[4], " ")[ifTrueLoc])]
		currentMonkey.ifFalse = &monkeys[parseToInt(strings.Split(monkeyInfo[5], " ")[ifFalseLoc])]
		divisor := parseToInt(strings.Split(monkeyInfo[3], " ")[divisorLoc])

		switch operator {
		case '+':
			if operandStr != "old" {
				operand := parseToInt(operandStr)
				currentMonkey.operation = func(i uint64) uint64 { return i + operand }
			} else {
				currentMonkey.operation = func(i uint64) uint64 { return i + i }
			}

		case '*':
			if operandStr != "old" {
				operand := parseToInt(operandStr)
				currentMonkey.operation = func(i uint64) uint64 { return i * operand }
			} else {
				currentMonkey.operation = func(i uint64) uint64 { return i * i }
			}
		default:
			log.Fatalf("Unknown operator: %s", string(operator))
		}

		currentMonkey.divisor = int(divisor)
	}

	return monkeys
}

func businessFactor(monkeys []Monkey) int {
	first, second := 0, 0
	for _, monkey := range monkeys {
		if monkey.itemsInspected > first {
			second = first
			first = monkey.itemsInspected
		} else if monkey.itemsInspected > second && monkey.itemsInspected < first {
			second = monkey.itemsInspected
		}
	}

	return first * second
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		gcdVal := gcd(int(result), nums[i])
		result = (result * num) / gcdVal
	}

	return result
}

func playRounds(rounds int, worryFactor int, monkeys *[]Monkey) {
	divisors := make([]int, 0, len(*monkeys))
	for _, v := range *monkeys {
		divisors = append(divisors, v.divisor)
	}

	optimizationFactor := lcm(divisors...)

	for i := 0; i < rounds; i++ {
		for j := 0; j < len(*monkeys); j++ {
			(*monkeys)[j].inspectItems(worryFactor, optimizationFactor)
		}
	}
}

func parseToInt(inp string) uint64 {
	res, err := strconv.ParseUint(inp, 10, 64)
	if err != nil {
		log.Fatalf("Trouble converting to int: %s", inp)
	}
	return res
}
