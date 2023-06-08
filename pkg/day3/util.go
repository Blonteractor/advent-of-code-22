package day3

const (
	asciiUpperCaseStart = 65
	asciiLowerCaseStart = 97
)

type Set map[rune]struct{}

func (s *Set) Add(item rune) {
	(*s)[item] = struct{}{}
}

func SetFromString(input string) Set {
	set := make(Set)
	for _, v := range input {
		set.Add(v)
	}
	return set
}

func (s *Set) Intersection(other Set) Set {
	intersection := make(Set)
	if len(*s) >= len(other) {
		for item := range other {
			if _, contains := (*s)[item]; contains {
				intersection.Add(item)
			}
		}
	} else {
		for item := range *s {
			if _, contains := other[item]; contains {
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
