package util

type HashSet[T comparable] map[T]struct{}

func (s *HashSet[T]) Add(item T) bool {
	_, exists := (*s)[item]
	if exists {
		return false
	}
	(*s)[item] = struct{}{}
	return true
}

func (s *HashSet[T]) Remove(item T) bool {
	_, exists := (*s)[item]
	if exists {
		return false
	}
	delete(*s, item)
	return true
}

func (s HashSet[T]) Contains(item T) bool {
	_, exists := (s)[item]
	return exists
}

func (s HashSet[T]) Len() int {
	return len(s)
}

func SetFromString(input string) HashSet[rune] {
	set := make(HashSet[rune])
	for _, v := range input {
		set.Add(v)
	}
	return set
}
