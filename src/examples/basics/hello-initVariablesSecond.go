package main

/**
 * The program demonstrates that go can infer the type of values
 *
 * author Gaurav Gaur
 **/
 
import (
	"fmt"
	"reflect"
)

var (
	name = "Gaurav"
	course = "Fundamentals of Go Programming"
	module = 3.2
)

func main() {
	fmt.Println("name is set to ", name, " and is of type ", reflect.TypeOf(name))
	fmt.Println("module is set to ", module, " and is of type ", reflect.TypeOf(module))
}