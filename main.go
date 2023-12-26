package main

import (
	"fmt"
	"regexp"
)

func main() {
	// compile ur regexp
	// re := regexp.MustCompile(`[a-c]+\d*`)
	// num_regexp_1 := regexp.MustCompile(`\(\d\d\d\) \d\d\d-\d\d\d\d`)
	// num_regexp_2 := regexp.MustCompile(`\(\d{3}\) \d{3}-\d{4}`) // simplified using braced quantifiers

	// // use this to see if there's a match
	// fmt.Println(num_regexp_1.MatchString("(916) 582-0183"))
	// fmt.Println(num_regexp_2.MatchString("(916) 582-0183"))

	// // use this to find to find the string that matches the regular expression
	// found_str := re.FindString("123abc332!")
	// fmt.Println(found_str)

	// temp
	// rr := regexp.MustCompile(`User \s*\b[a-zA-Z]+\d+\b`)
	u_regexp := regexp.MustCompile(`\b[a-zA-Z]+\d+\b`)
	str := "[WRN] User James123 has exceeded storage space."
	// user_name := u_regexp.FindString(str)
	fmt.Println(u_regexp.ReplaceAllString(str, ""))
}
