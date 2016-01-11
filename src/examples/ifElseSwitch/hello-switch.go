package main

/**
 * The program displays the basic usage and syntax of swith statment in Go
 *
 * author Gaurav Gaur
 *
 **/

import (
 	"fmt"
)

func main() {

	switch "Gaurav" {
		case "Something": fmt.Println("Is that even a name")
		case "Bob": fmt.Println("That is not your name")
		case "Gaurav": fmt.Println("You don't need switch statement to know your name")
		default: fmt.Println("Try to find your name, I hope you know it!")
	}
}