package ranger

import (
	"errors"
)

var (
	errInvalidStart = errors.New("start must be less than end")
)

type Range[T number] struct {
	start T
	end   T
}

func NewRange[T number](start T, end T) (*Range[T], error) {
	if start > end {
		return nil, errInvalidStart
	}

	return &Range[T]{
		start: start,
		end:   end,
	}, nil
}

func (r *Range[T]) Start() T { return r.start }

func (r *Range[T]) End() T { return r.end }
