package main

/**
 * The program displays the use of range looping and 
 * nested for loops
 *
 * @author: Gaurav Gaur
 * 
 **/

import (
	"fmt"
)

func main() {

	// This is a slice which is unordered list of items
	myList := []string{"One", "Two", "Three", "Four"}
	anotherList := []string{"A", "B", "C", "D"}

	for _, num := range myList {
		fmt.Println(num)
		for _, alpha := range anotherList {
			fmt.Println(alpha)
		}
	}
}