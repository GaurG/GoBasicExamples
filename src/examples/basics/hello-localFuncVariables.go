package main

/**
 * The program displays the use of println and reflect funtion
 * and how Go can infer the type of variable
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

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
}