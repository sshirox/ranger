package ranger

func (s *Sequence[T]) All(fn func(T) bool) bool {
	for _, v := range s.values {
		if !fn(v) {
			return false
		}
	}

	return true
}
