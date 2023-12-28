package using_stringers

import "fmt"

type Person struct {
	name string
	age  int
}

// when u call print on a method, it will check for the
// implementation of the Stringers interface and display that
func (p Person) String() string {
	return fmt.Sprintf("Name is: %s, and age is: %d", p.name, p.age)
}

func Run() {
	p := Person{}
	p.name = "Nnaemeka"
	p.age = 22
	fmt.Println(p) // Name is: Nnaemeka, and age is: 22
}
