package main

/**
 * The program demonstrates the go can infer the type of values and assign it
 * to appropriate variables
 *
 * author Gaurav Gaur
 **/

import (
	"fmt"
	"reflect"
)

var (
	name, course, module = "Gaurav", "Fundamentals of Go Programming", 3.2
)

func main() {
	fmt.Println("name is set to ", name, " and is of type ", reflect.TypeOf(name))
	fmt.Println("module is set to ", module, " and is of type ", reflect.TypeOf(module))
}