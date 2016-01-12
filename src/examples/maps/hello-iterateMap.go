package main

/**
 * The program illustrates how we can itereate over the maps.
 *
 * @author Gaurav Gaur
 *
 **/
import (
	"fmt"
)

func main() {

	myMap := map[string]int {"A":1, "B":2,"C":3,"D":4, "E":5,"F":6}

	for key, value := range myMap {
		fmt.Printf("\nKey is: %v, value is: %v", key, value)
	}

}