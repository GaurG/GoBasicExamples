package main

/**
 * The program displays the usage of local variable. 
 * Make a note that how these are assigned the default values
 *
 * @author Gaurav Gaur
 *
 **/

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.0000
	b := 3

	fmt.Println("\n A is type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b) )

	//c := a + b the statement will result in mismatch exception
	c := int(a) + b

	fmt.Println("\n C has a value:", c, "and is of type:", reflect.TypeOf(c))
}