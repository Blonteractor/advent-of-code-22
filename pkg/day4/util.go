package day4

import (
	"log"
	"strconv"
	"strings"
)

type NumberRange struct {
	lower int
	upper int
}

func numberRangeFromString(input string) NumberRange {
	s := strings.Split(input, "-")

	a, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatalf("Trouble parsing input: %s", input)
	}

	b, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatalf("Trouble parsing input: %s", input)
	}

	return NumberRange{
		lower: a,
		upper: b,
	}
}

func (n NumberRange) isContainedIn(other NumberRange) bool {
	return n.lower >= other.lower && n.upper <= other.upper
}

func (n NumberRange) isOverlapping(other NumberRange) bool {
	if n.lower >= other.lower {
		return n.lower <= other.upper
	} else {
		return n.upper >= other.lower
	}
}
