package utils

type Set[T comparable] map[T]struct{}

func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) AddAll(items []T) {
	for _, item := range items {
		s.Add(item)
	}
}
