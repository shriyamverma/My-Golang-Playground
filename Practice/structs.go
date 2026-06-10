package main

import "fmt"

type Employee struct {
	name     string
	age      int
	isRemote bool
	address  Address // composition
}

type Address struct {
	street string
	city   string
}

func (e Employee) updateNameByValue(name string) {
	e.name = name
}

func (e *Employee) updateNameByRef(name string) {
	e.name = name
}

func (a Address) printAddress() {
	fmt.Printf("Street: %s, City: %s\n", a.street, a.city)
}

func main() {
	address := Address{
		street: "123 Main St",
		city:   "New York",
	}
	employee1 := Employee{
		name:     "John",
		age:      42,
		isRemote: true,
		address:  address,
	}

	fmt.Println("Employee Name: ", employee1.name)
	fmt.Println("Employee Age: ", employee1.age)
	fmt.Println("Employee IsRemote: ", employee1.isRemote)
	employee1.address.printAddress()

	fmt.Println("-------------------------")
	// Update name (Wrong Way)
	employee1.updateNameByValue("Bob")
	fmt.Println("[UpdateName ❌] Employee Name: ", employee1.name)

	// Update name (Correct Way)
	employee1.updateNameByRef("Bob")
	fmt.Println("[UpdateName ✅] Employee Name: ", employee1.name)
	fmt.Println("-------------------------")

	// anonymous structs
	job := struct {
		title  string
		salary int
	}{
		title:  "Software Engineer",
		salary: 100000,
	}

	fmt.Println("Job Title: ", job.title)

	// To mutate the values directly
	employeePtr := &employee1
	employeePtr.age = 31
	fmt.Println("Employee Age after mutation: ", employee1.age)

}
