package ranger

import (
    "errors"
    "math"
)

var (
    errZeroStep         = errors.New("step must not be zero")
    errInvalidDirection = errors.New("step direction does not match range direction")
)

type Sequence[T number] struct {
    values []T
}

func NewSequence[T number](values []T) *Sequence[T] {
    return &Sequence[T]{values: values}
}

func (s *Sequence[T]) Values() []T {
    res := make([]T, len(s.values))

    copy(res, s.values)
    
    return res
}

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
    values := make([]T, 0, capacity)

    if step > 0 {
        for v := start; v <= end; v += step {
            values = append(values, v)
        }
    } else {
        for v := start; v >= end; v += step {
            values = append(values, v)
        }
    }

    return &Sequence[T]{values: values}, nil
}
