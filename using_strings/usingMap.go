package using_strings

import (
	"fmt"
	"strings"
)

func UseMapFunc() {
	var s string = "Nnaemeka"
	var mapFunc func(r rune) rune = func(r rune) rune {
		if r == 'e' {
			// this essentially drops e from the string
			return -1
		}
		return r
	}
	var newS string = strings.Map(mapFunc, s)
	fmt.Println("old name is", s, "and new name is", newS)
}
