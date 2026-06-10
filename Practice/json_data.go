package main

import (
	"encoding/json"
	"fmt"
)

type Employee1 struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	IsRemote bool     `json:"isRemote"`
	Address  Address1 `json:"address"`
}

type Address1 struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

func (e Employee1) updateNameByValue(name string) {
	e.Name = name
}

func (e *Employee1) updateNameByRef(name string) {
	e.Name = name
}

func (a Address1) printAddress() {
	fmt.Printf("Street: %s, City: %s\n", a.Street, a.City)
}

func main() {
	address := Address1{
		Street: "123 Main St",
		City:   "New York",
	}
	employee1 := Employee1{
		Name:     "John",
		Age:      42,
		IsRemote: true,
		Address:  address,
	}

	fmt.Println("Employee1 Name: ", employee1.Name)
	fmt.Println("Employee1 Age: ", employee1.Age)
	fmt.Println("Employee1 IsRemote: ", employee1.IsRemote)
	employee1.Address.printAddress()

	jsonData, _ := json.Marshal(employee1)
	fmt.Println(jsonData) // bytes
	fmt.Println(string(jsonData))
}
