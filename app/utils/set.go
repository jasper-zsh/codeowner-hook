package utils

type Set[T string] map[T]struct{}

func (s Set[T]) Add(item ...T) {
	for _, i := range item {
		s[i] = struct{}{}
	}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) ToArray() []T {
	arr := make([]T, 0, len(s))
	for i := range s {
		arr = append(arr, i)
	}
	return arr
}
