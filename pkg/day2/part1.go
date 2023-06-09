package day2

import (
	"fmt"
	"strings"
)

func SolvePart1(input string) string {
	strategyGuide := strings.Split(input, "\n")

	score := 0
	for _, strat := range strategyGuide {
		var myMove Move
		var opponentMove Move

		switch strat[0] {
		case 'A':
			opponentMove = Rock
		case 'B':
			opponentMove = Paper
		case 'C':
			opponentMove = Scissors
		}

		switch strat[2] {
		case 'X':
			myMove = Rock
		case 'Y':
			myMove = Paper
		case 'Z':
			myMove = Scissors
		}

		result := myMove.play(opponentMove)

		score += int(myMove) + int(result)
	}

	return fmt.Sprint(score)
}
