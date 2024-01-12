package testing

import "golang.org/x/exp/constraints"

type numeric interface {
	constraints.Integer | constraints.Float
}

func Sum[T numeric](numbers ...T) T {
	var s T
	for _, num := range numbers {
		s += num
	}
	return s
}
