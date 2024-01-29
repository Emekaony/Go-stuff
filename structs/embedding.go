package structs

import "fmt"

type Human struct {
	firstName string
	lastName  string
	age       int
}

type SoftwareEngineer struct {
	Human
	Role    string
	Company string
}

func Run() {
	metaEngineer := SoftwareEngineer{
		Human: Human{
			firstName: "Nnaemeka",
			lastName:  "Onyeokoro",
			age:       22,
		},
		Role:    "Backend Engineer",
		Company: "Meta",
	}
	fmt.Println(metaEngineer.firstName) // works because of struct embedding
}
