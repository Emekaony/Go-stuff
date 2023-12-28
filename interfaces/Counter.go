package interfaces

// look out for empty interfaces in go

import "fmt"

type Counter interface {
	Add(increment int)
	Value() int
}

type Stats struct {
	value int
}

// any type that has the required methods in an interface
// implicitly implements that interface. There is no `implements` keyword in go
func (s Stats) Add(value int) {
	s.value += value
}

func (s Stats) Value() int {
	return s.value
}

func (s Stats) sayHello(name string) {
	message := fmt.Sprintf("hello there, %s", name)
	fmt.Println(message)
}
