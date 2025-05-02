package ranger

func (s *Sequence[T]) Map(fn func(T) T) *Sequence[T] {
	result := make([]T, len(s.values))

	for i, value := range s.values {
		result[i] = fn(value)
	}

	return &Sequence[T]{values: result}
}
