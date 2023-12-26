package main

import (
	"fmt"
	"fundamentals/structs"
)

func main() {
	person1 := structs.NewPerson("Nnaemeka", "Onyeokoro", 22)
	fmt.Printf("Your name is %s\n", person1.GetFirstName())
}
