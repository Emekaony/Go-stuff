package async

import (
	"fmt"
	"time"
)

// goroutines are lightweight threads
// they run concurrently together

func sayHello(name string) {
	fmt.Println("hello there", name)
}

func Run() {
	// regular call to sayHello(name stering)
	sayHello("Nnaemeka")

	// now let us use goroutines
	go sayHello("kamsi")
	go func(msg string) {
		fmt.Println("sent in", msg, "to async function")
	}("dummy text")

	time.Sleep(time.Second) // sleep for 1 second while the
}
