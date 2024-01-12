package benchmark

import "golang.org/x/exp/constraints"

type numeric interface {
	constraints.Integer | constraints.Float
}

// Generic sum function
func Sum[T numeric](numbers ...T) T {
	var s T
	for _, num := range numbers {
		s += num
	}
	return s
}

// sum for integers
func SumInt(nums ...int) int {
	var s int
	for _, num := range nums {
		s += num
	}
	return s
}
