package main

import "fmt"

func main() {
	arr := []int{}
	arr2 := make([]int, len(arr)) // so we make a slice with the same length as the first array
	arr2 = append(arr2, []int{88, 99, 100}...)
	num := copy(arr2, arr) // now we copy the stuff from arr into arr2
	fmt.Println(arr2, num)
}
