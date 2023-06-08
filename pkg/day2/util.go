package day2

type Move int;
const (
	Rock Move = 1
	Paper Move = 2
	Scissors Move = 3
)

type Result int;
const (
	Loss Result = 0
	Draw Result = 3
	Win Result = 6
)

func(a Move) play(b Move) Result {
	if a == b {
		return Draw
	} else if (a == Rock && b == Scissors) || (a == Scissors && b == Paper) || (a == Paper && b == Rock) {
		return Win
	} else {
		return Loss	
	}
}
