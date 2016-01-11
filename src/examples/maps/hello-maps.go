package main

/**
 * The program illustrates how maps can be declared in Go.
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