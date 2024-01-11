package main

import (
	"fmt"
)

func main() {
	var i any = 's'

	switch v := i.(type) {
	case int:
		fmt.Println(i, "is an int")
		fmt.Println(i.(int))
	case string:
		fmt.Println(v, " is a string")
	case rune:
		fmt.Println(string(v), "is a rune")
		fmt.Println(string(v))
	default:
		fmt.Println(i, "is a not a string, int, or rune!")
		fmt.Println("it is not a string or an int")
	}
}
