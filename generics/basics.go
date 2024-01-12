package generics

import "fmt"

func sum[T int | float64](a, b T) T {
	return a + b
}

func Run() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(4.4, 3.9))
}
