package ranger

func (r *Range[T]) ToSlice(step T) []T {
	if step == 0 {
		panic("step must not be zero")
	}

	start, end := r.Start(), r.End()

	if (start < end && step < 0) || (start > end && step > 0) {
		panic("step direction does not match range direction")
	}

	var result []T

	if step > 0 {
		for v := start; v <= end; v += step {
			result = append(result, v)
		}
	} else {
		for v := start; v >= end; v += step {
			result = append(result, v)
		}
	}

	return result
}
