package day12

import (
	"fmt"
	"strings"
)

const (
	startSymbol = 'S'
	endSymbol   = 'E'
)

func SolvePart1(input string) string {
	lines := strings.Split(input, "\n")
	heatmap := make([][]*Node, 0, len(lines))
	var start *Node
	var end *Node

	for i, v := range lines {
		line := make([]*Node, 0, len(v))
		for j, v2 := range v {
			node := node(v2, j, i)
			line = append(line, &node)
			if v2 == startSymbol {
				start = &node
				start.height = 'a'
			} else if v2 == endSymbol {
				end = &node
				end.isEnd = true
				end.height = 'z'
			}
		}
		heatmap = append(heatmap, line)
	}

	steps := aStar(start, end, heatmap)
	
	return fmt.Sprint(steps)
}
