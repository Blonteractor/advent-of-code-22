package day3

import "github.com/Bonteractor/advent-of-code-22/pkg/util"

const (
	asciiUpperCaseStart = 65
	asciiLowerCaseStart = 97
)

func intersection(s util.HashSet[rune], other util.HashSet[rune]) util.HashSet[rune] {
	intersection := make(util.HashSet[rune])
	if len(s) >= len(other) {
		for item := range other {
			if s.Contains(item) {
				intersection.Add(item)
			}
		}
	} else {
		for item := range s {
			if s.Contains(item) {
				intersection.Add(item)
			}
		}
	}

	return intersection
}

func CalcPriority(char rune) int {
	if char >= asciiLowerCaseStart {
		return int(char - asciiLowerCaseStart + 1)
	} else {
		return int(char - asciiUpperCaseStart + 27)
	}
}
