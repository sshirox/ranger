package ranger

func (r *Range[T]) Contains(value T) bool {
	return value >= r.Start() && value <= r.End()
}
