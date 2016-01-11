package main

/**
 * The program illustrates on how we create slice.
 *
 * @author Gaurav Gaur
 *
 **/

import (
	"fmt"
)

func main() {
	// Declare a slice with type, len and capacity
	mySlice := make([]string, 5, 10)
	fmt.Printf("Length is: %d \n capacity is: %d\n\n", len(mySlice), cap(mySlice))

	//Declare a slice with default capacity
	mySecondSlice := make([]string, 5)
	fmt.Printf("Length is: %d \n capacity is: %d\n\n", len(mySecondSlice), cap(mySecondSlice))

	//Declare a slice with initialized value
	myThirdSlice := []string{"value1", "value2", "value3"}
	fmt.Printf("Length is: %d \n capacity is: %d\n\n", len(myThirdSlice), cap(myThirdSlice))

	fmt.Println(myThirdSlice[1])
	myThirdSlice[1] = "third"
	fmt.Println(myThirdSlice[1])


}