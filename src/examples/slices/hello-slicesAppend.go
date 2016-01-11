package main

/**
 * The program illustarates append functionality of slices
 *
 * @author Gaurav Gaur
 *
 **/

import (
	"fmt"
)

func main() {
	// Declare a slice with type, len and capacity
	mySlice := []int{1,2,3,4}

	for _, i := range mySlice {
		fmt.Println("for range loop", i)
	}

	secondSlice := []int{10,20,30,50}

	// ... shows that we are appending the values of second slice
	mySlice = append(mySlice, secondSlice...)
	fmt.Println(mySlice)
	


}