package ranger

func (r *Range[T]) Equals(other *Range[T]) bool {
	if r == nil || other == nil {
		return false
	}

	return r.Start() == other.Start() && r.End() == other.End()
}
