package structs

import (
	"encoding/json"
	"fmt"
)

type Personn struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	Personn
	ManagerID int
}

type Contractor struct {
	Personn
	CompanyID int
}

func EncodeJson() {
	employees := []Employee{
		{
			Personn: Personn{
				LastName: "Doe", FirstName: "John", ID: 22445, Address: "John street, 11th avenue",
			},
		},
		{
			Personn: Personn{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	json.Unmarshal(data, &decoded)
	// fmt.Println(decoded[0].FirstName)
	fmt.Printf("%v\n", decoded)
}
