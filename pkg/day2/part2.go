package day2

import (
	"strings"
)

func SolvePart2(input string) int {
	strategyGuide := strings.Split(input, "\n")
	
	score := 0
	for _, strat := range strategyGuide {
		var opponentMove Move;
		var myMove Move;
		var exptectedResult Result;
		
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
			exptectedResult = Loss
		case 'Y':
			exptectedResult = Draw
		case 'Z':
			exptectedResult = Win			
		}	
		
		if exptectedResult == Draw {
			myMove = opponentMove
		} else if exptectedResult == Win {
			idx := int(opponentMove) + 1 
			if idx == 4 {
				idx = 1
			}
			myMove = Move(idx)
		} else if exptectedResult == Loss {
			idx := int(opponentMove) - 1 
			if idx == 0 {
				idx = 3
			}
			myMove = Move(idx)
		}
		
		score += int(myMove) + int(exptectedResult)
	}
	
	return score
}