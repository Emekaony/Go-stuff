package loops

import (
	"fmt"
)

func Run() {
	s := "Nnaemeka"
	for _, char := range s {
		fmt.Println("char is ", string(char))
	}
}
