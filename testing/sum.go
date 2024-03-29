package testing

import "golang.org/x/exp/constraints"

type numeric interface {
	// this is similar to the ~(tilde) thing where any type that has an int or float works
	constraints.Integer | constraints.Float
}

func Sum[T numeric](numbers ...T) T {
	var s T
	for _, num := range numbers {
		s += num
	}
	return s
}
