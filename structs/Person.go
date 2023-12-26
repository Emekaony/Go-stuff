package structs

type Person struct {
	firstname string
	lastname  string
	age       int
}

func (p *Person) GetAge() int {
	return p.age
}

// this one should be a pointer receier method because it modifies
// the struct that was passed into it!
func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) GetFirstName() string {
	return p.firstname
}

// this is how u implement the constructor method for the Person type
func NewPerson(firstname, lastname string, age int) *Person {
	return &Person{
		age:       age,
		firstname: firstname,
		lastname:  lastname,
	}
}
