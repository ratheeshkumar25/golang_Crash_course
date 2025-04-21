package main

import "fmt"

// interfaces in Go are considered a data type but with a very specific purpose. They are used to define a set of method signatures that a type must implement
// declaring interface using type keyword

//we have company interface that contain employee method and its type

type Company interface {
	Employees() string
	Department() string
}

// we have to create struct that satisfies the interfaces
type Staff struct {
	Name        string
	Departments string
}

// Now we need to impement the Employee method signature that will return employees name
func (e *Staff) Employees() string {
	return e.Name
}

// Now we need to implement the Deparment method
func (e *Staff) Department() string {
	return e.Departments
}

func main() {

	// Create a Staff custom data type.
	// Use a pointer to match the method receiver
	staff := &Staff{Name: "Ratheesh", Departments: "Go Developer"}
	var company Company

	// Assign the Staff instance to the Company interface.
	company = staff
	fmt.Println("Employee Name", company.Employees())
	fmt.Println("Employee Name", company.Department())
}
