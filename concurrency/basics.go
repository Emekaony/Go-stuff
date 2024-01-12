package concurrency

import "time"

func sum(numbers ...int) {
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	println(sum)
}

func Run() {
	go sum([]int{1, 2, 3, 4}...)
	time.Sleep(100 * time.Millisecond) // let the main thread sleep for a hundres milliseconds
}
