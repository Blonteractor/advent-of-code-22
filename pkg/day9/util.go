package day9

type Set map[Position]struct{}

func (s *Set) Add(item Position) {
	(*s)[item] = struct{}{}
}

type Position struct {
	x int
	y int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func signum(n int) int {
	if n == 0 {
		return 0
	} else if n > 0 {
		return 1
	} else {
		return -1
	}
}

func (p Position) isTouching(other Position) bool {
	if p.x == other.x && p.y == other.y {
		return true
	}

	// Check if the points are touching diagonally
	if abs(p.x-other.x) == 1 && abs(p.y-other.y) == 1 {
		return true
	}

	// Check if the points are touching in the up, down, left, or right direction
	if (p.x == other.x && abs(p.y-other.y) == 1) || (p.y == other.y && abs(p.x-other.x) == 1) {
		return true
	}

	// Points are not touching diagonally or in the up, down, left, or right direction
	return false
}

func (p Position) moveTail(tail *Position) {
	if tail != nil && !p.isTouching(*tail) {
		x_offset := p.x - tail.x
		y_offset := p.y - tail.y

		tail.x += signum(x_offset)
		tail.y += signum(y_offset)
	}
}

func (p *Position) moveBy(offset Position) {
	p.x += offset.x
	p.y += offset.y
}
