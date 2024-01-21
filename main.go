package main

import "fmt"

func main() {
	// s := "emeka"
	// fmt.Println(rune(s[0]))
	j := []byte{101, 109, 101, 107, 97} // this is a byte representing the unicode code point for each rune.
	fmt.Println(string(j))
}
