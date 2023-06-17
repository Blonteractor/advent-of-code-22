package day10

import (
	"fmt"
	"os"
	"testing"
)

const (
	part1Ans = 13140
	part2Ans = ""
)

func Test_SolvePart1(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	ans := SolvePart1(string(input))

	if ans != fmt.Sprint(part1Ans) {
		t.Errorf("Expected %d, got %s", part1Ans, ans)
	}
}

func Test_SolvePart2(t *testing.T) {
	input, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	ans := SolvePart2(string(input))

	if ans != fmt.Sprint(part2Ans) {
		t.Errorf("Expected %s, got %s", part2Ans, ans)
	}
}
