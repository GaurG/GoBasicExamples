package main

/**
 * The program illustrates that how we can delcare 
 * structs in Go
 *
 * @author Gaurav Gaur
 *
 **/

import (
	"fmt"
)

func main() {

	type employeeDetails struct {
		Name string
		Manager string
		Salary float64
	}

	// other ways of initializing a struct
	//var GauravGaur employeeDetails
	//GauravGaur := new(employeeDetails)

	// observe how we can assign a value if know the name
	// of variable
	GauravGaur := employeeDetails {
		Manager: "Some one",
		Name: "Gaurav Gaur",
		Salary: 100,
	}

	fmt.Println(GauravGaur)

	// accessing single value
	fmt.Println("manager name is: ", GauravGaur.Manager)

	// change the value of field
	GauravGaur.Manager = "Someone Else"
	fmt.Println(GauravGaur)


}