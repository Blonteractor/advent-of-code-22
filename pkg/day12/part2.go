package day12

import (
	"fmt"
	"strings"
)

func SolvePart2(input string) string {
	lines := strings.Split(input, "\n")
	heatmap := make([][]*Node, 0, len(lines))
	start := make([]*Node, 0)
	var end *Node

	for i, v := range lines {
		line := make([]*Node, 0, len(v))
		for j, v2 := range v {
			node := node(v2, j, i)
			line = append(line, &node)
			if v2 == startSymbol || v2 == 'a' {
				node.height = 'a'
				start = append(start, &node)
			} else if v2 == endSymbol {
				end = &node
				end.isEnd = true
				end.height = 'z'
			}
		}
		heatmap = append(heatmap, line)
	}
	
	min := aStar(start[0], end, heatmap)
	for i := 1; i < len(start); i++ {
		steps := aStar(start[i], end, heatmap)
		if steps < min {
			min = steps
		}
	}
	
	return fmt.Sprint(min)
}
