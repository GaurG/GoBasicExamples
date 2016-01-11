package main

/**
 * The program illustarates that the capacity of under lying array
 * increases as we keep adding values.
 *
 * @author Gaurav Gaur
 *
 **/

import (
	"fmt"
)

func main() {
	// Declare a slice with type, len and capacity
	mySlice := make([]int, 1, 4)

	fmt.Printf("length is %d \nand capacity is %d \n\n", len(mySlice), cap(mySlice))

	for i := 1; i<17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("capacity is %d \n", cap(mySlice))		
	}
	


}