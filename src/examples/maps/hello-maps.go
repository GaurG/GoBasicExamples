package main

/**
 * The program illustrates how maps can be declared in Go.
 * Maps are:
 * 1. Passed to functions by reference similar to java
 * 2. Unsafe for concurrency
 * 3. Good practice to specify the size of the map
 *
 * @author Gaurav Gaur
 *
 **/
import (
	"fmt"
)

func main() {

	// map[keys]values
	myIplWonMap := make(map[string]int)

	myIplWonMap["Chennai"] = 6
	myIplWonMap["Delhi"] = 1
	myIplWonMap["Pune"] = 0

	fmt.Println("my Map ", myIplWonMap)

	anotherIplMap := map[string]int {
		"Chennai":6,
		"Delhi":1,
		"Pune":0,
	}

	fmt.Println("my another Map ", anotherIplMap)

}