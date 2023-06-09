package day7

import (
	"log"
	"strconv"
	"strings"
)

// Dirs have size 0, Files are leaf nodes
type Node struct {
	name     string
	children []*Node
	parent   *Node
	size     int
}

func (node *Node) Directories() []*Node {
	var flattened []*Node
	stack := []*Node{node}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		flattened = append(flattened, current)

		for i := len(current.children) - 1; i >= 0; i-- {
			child := current.children[i]
			if !child.isFile() {
				stack = append(stack, current.children[i])
			}
		}
	}

	return flattened
}

func (n *Node) EmplaceDirectory(name string, children []*Node) *Node {
	node := Node{
		name:     name,
		children: children,
		parent:   n,
		size:     0,
	}
	n.children = append(n.children, &node)

	return &node
}

func (n *Node) EmplaceFile(name string, size int) *Node {
	node := Node{
		name:     name,
		children: nil,
		parent:   n,
		size:     size,
	}
	n.children = append(n.children, &node)

	return &node
}

func (n Node) isFile() bool {
	return len(n.children) == 0 && n.size != 0
}

func (n Node) Root() Node {
	current := n
	for current.name != "/" {
		current = *current.parent
	}
	return current
}

func (n Node) CalculateSize() int {
	totalSize := 0

	for _, v := range n.children {
		if v.isFile() {
			totalSize += v.size
		} else {
			totalSize += v.CalculateSize()
		}
	}

	return totalSize
}

func (n Node) Walk(ident string) (*Node, bool) {
	for _, v := range n.children {
		if v.name == ident {
			return v, true
		}
	}

	return nil, false
}

func parseInput(input string) Node {
	currentDir := &Node{
		name:     "start",
		parent:   nil,
		children: nil,
		size:     0,
	}
	for _, v := range strings.Split(input, "\n") {
		if v[0] == '$' {
			prompt := strings.Split(v, " ")
			command := prompt[1]
			if command == "cd" {
				arg := prompt[2]
				// If .., go back to prev dir, otherwise update prev dir
				if arg == ".." {
					currentDir = currentDir.parent
				} else {
					dir, exists := currentDir.Walk(arg)
					if exists {
						currentDir = dir
					} else {
						// If it doesnt exist, create it in the current folder with no children(for now)
						currentDir = currentDir.EmplaceDirectory(arg, nil)
					}
				}
			} else if command == "ls" {
				continue
			}
		} else {
			lsCommand := strings.Split(v, " ")

			// Can be file size or the string 'dir' indicating a directory
			metadata := lsCommand[0]
			ident := lsCommand[1]

			if metadata == "dir" {
				_, exists := currentDir.Walk(ident)
				if !exists {
					currentDir.EmplaceDirectory(ident, nil)
				}
			} else {
				size, err := strconv.Atoi(metadata)
				if err != nil {
					log.Fatalf("Trouble parsing ls output, file size: %s", err)
				}

				currentDir.EmplaceFile(ident, size)
			}
		}
	}

	return currentDir.Root()
}
