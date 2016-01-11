package main

/**
 * The program displays the usage of normal break and label break
 *
 * @author Gaurav Gaur
 *
 **/

import (
	"fmt"
)

func main() {

	// This is a slice which is unordered list of items
	myList := []string{"One", "Two", "Three", "Four"}
	anotherList := []string{"A", "B", "C", "D"}

	fmt.Println("usual break at B")

	for _, num := range myList {
		fmt.Print("\n",num)
		for _, alpha := range anotherList {
			fmt.Print("\t", alpha)
			if("C" == alpha) {
				break
			}
		}
	}


	fmt.Println("\n\nlabel break")
	
	bothLoops:
	for _, num := range myList {
		fmt.Print("\n",num)
		for _, alpha := range anotherList {
			fmt.Print("\t", alpha)
			if("C" == alpha && "Three" == num) {
				break bothLoops
			}
		}
	}
}