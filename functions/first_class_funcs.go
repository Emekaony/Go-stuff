package functions

import "fmt"

func FirstGreeting(name string) string {
	message := fmt.Sprintf("hello there, %s", name)
	return message
}

func SecondGreeting(name string) string {
	message := fmt.Sprintf("Hello %s, how are you doing?", name)
	return message
}

func Run() {
	name := "Nnaemeka"
	f1 := FirstGreeting(name)
	f2 := SecondGreeting(name)

	fmt.Println(f1, "\n", f2)
}
