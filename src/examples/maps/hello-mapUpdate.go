package main

/**
 * The program illustrates how we update map elements
 *
 * @author Gaurav Gaur
 *
 **/
import (
	"fmt"
)

func main() {

	myMap := map[string]int {"A":1, "B":2,"C":3,"D":4, "E":5,"F":6}

	fmt.Println("value of C is", myMap["C"])

	myMap["C"] = 300

	fmt.Println("updated value of C is", myMap["C"])

	myMap["G"] = 7

	fmt.Println("after adding a new value G", myMap)

	delete(myMap, "G")

	//bogus delete just to check what happens
	delete(myMap, "H")

	fmt.Println("after deleting a new value G", myMap)

}