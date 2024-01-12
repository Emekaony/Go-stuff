package interfaces

import "fmt"

// swap takes in two values and returns them swapped
// because every type imeplements any
// so we do not have to specify the type - will go further into this later.
func swap[T any](a, b T) (T, T) {
	return b, a
}

// interface constraints
type numeric interface {
	int | float64
}

func sum[T numeric](a, b T) T {
	return a + b
}

func Run() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(9.9, 55.32))
}
