package ranger

import (
	"errors"
	"math"
)

type Sequence[T number] struct {
	data []T
}

var (
	errZeroStep         = errors.New("step must not be zero")
	errInvalidDirection = errors.New("step direction does not match range direction")
)

func (r *Range[T]) ToSequence(step T) (*Sequence[T], error) {
	if step == 0 {
		return nil, errZeroStep
	}

	start, end := r.Start(), r.End()

	if (start < end && step < 0) || (start > end && step > 0) {
		return nil, errInvalidDirection
	}

	var count T
	if step > 0 {
		count = (end - start) / step
	} else {
		count = (start - end) / (-step)
	}

	capacity := int(math.Ceil(float64(count))) + 1
	result := make([]T, 0, capacity)

	if step > 0 {
		for v := start; v <= end; v += step {
			result = append(result, v)
		}
	} else {
		for v := start; v >= end; v += step {
			result = append(result, v)
		}
	}

	return &Sequence[T]{data: result}, nil
}
