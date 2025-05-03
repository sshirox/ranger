package ranger

func (s *Sequence[T]) Filter(fn func(T) bool) *Sequence[T] {
	result := make([]T, 0, len(s.values))

	for _, value := range s.values {
		if fn(value) {
			result = append(result, value)
		}
	}

	return &Sequence[T]{values: result}
}
