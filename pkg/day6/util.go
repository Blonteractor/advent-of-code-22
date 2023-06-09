package day6

import (
	"fmt"
)

func findFirstUniqueSet(input string, size int) (int, error) {
	i := size - 1
	for i < len(input) {
		var possibleMarker = make([]rune, 0)
		for j := size - 1; j >= 0; j-- {
			possibleMarker = append(possibleMarker, rune(input[i-j]))
		}

		duplicate := size
		for i, v := range possibleMarker {
			for j := i + 1; j < size; j++ {
				if v == possibleMarker[j] {
					duplicate = i
				}
			}
		}
		if duplicate == size {
			return i + 1, nil
		} else {
			i += duplicate + 1
		}

	}
	return 0, fmt.Errorf("bad input, no marker found")
}
