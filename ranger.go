package ranger

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	errInvalidBegin = errors.New("begin must be less than end")
)

type number interface {
	constraints.Integer | constraints.Float
}

type Range[T number] struct {
	begin T
	end   T
}

func NewRange[T number](begin T, end T) (*Range[T], error) {
	if begin > end {
		return nil, errInvalidBegin
	}

	return &Range[T]{
		begin: begin,
		end:   end,
	}, nil
}

func (r Range[T]) Begin() T {
	return r.begin
}

func (r Range[T]) End() T {
	return r.end
}
