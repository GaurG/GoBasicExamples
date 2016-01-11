package main

/**
 * The program introduced pointers and shows the usage of
 * '&' referencing and '*' de-referencing a pointer
 *
 * @author Gaurav Gaur
 *
 **/
import (
	"fmt"
	"reflect"
)

func main() {
	
	name := "Gaurav"
	//course := "Go fundamentals"
	module := "3.2"
	ptr := &module

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Memory address of *module* variable is ", ptr,
		"and the value of *module* is ", *ptr)
}