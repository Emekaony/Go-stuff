package concurrency

import (
	"fmt"
)

func sum(numbers []int, ch chan int) {
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	ch <- sum
}

func Run() {
	nums := []int{1, 2, 3, 4}

	// we must use the make built-in for channel initialization
	ch := make(chan int)

	go sum(nums, ch)

	// this is blocking
	result := <-ch

	fmt.Println("Result is:", result)
}
