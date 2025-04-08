package ranger

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}
