package main

/**
 * The program illustrates how we create multiple smaller slices
 * from a given slice.
 *
 * @author Gaurav Gaur
 *
 **/

import (
	"fmt"
)

func main() {

	mySlice := []int{1,2,3,4,5,6,7,8,9,10}

	// slice includes the first index and excludes the second index.
	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)

	// missing value is inferred as 0
	secondSlice := mySlice[:5]
	fmt.Println(secondSlice)

	// missing value is inferred as length of array - 1
	thirdSlice := mySlice[6:]
	fmt.Println(thirdSlice)
}