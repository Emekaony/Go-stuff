package type_conversions

import (
	"fmt"
	"reflect"
	"strconv"
)

func Example() {
	// define int and float64 values:
	const x int = 22
	const c float64 = float64(x)

	// do not forget that u can use the %T format specifier
	// to show the type of a variable
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("c is of type %T\n", c)

	// go from string to int
	var intString string = "42"
	var i, err = strconv.Atoi(intString)

	if err == nil {
		fmt.Printf("Successifully converted %s from %T to %d of type %T\n", intString, intString, i, i)
	} else {
		fmt.Printf("could not perform conversion due to error\n")
	}

	// go from int to string
	var number int = 12
	var s string = strconv.Itoa(number)
	fmt.Printf("converted %d from type %T to type %T\n", number, number, s)

	// get the type of a variable using reflect.TypeOf(value)
	var ss string = "emeka"
	var t uint = 22
	fmt.Printf("The type of %v is %s\n", ss, reflect.TypeOf(s))
	fmt.Printf("The type of %v is %s\n", t, reflect.TypeOf(t))
}
