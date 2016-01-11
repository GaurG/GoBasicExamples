package main

/**
 * The program uses pass by reference in GO
 *
 * @author Gaurav Gaur
 *
 **/
import (
	"fmt"
)

func main() {
	
	name := "Gaurav"
	course := "Go fundamentals"

	fmt.Println("\nHi", name, "you are currently watching", course)
	changeCourse(&course)
	fmt.Println("\nYou are now watching the course", course)
	
}

func changeCourse(course *string) {
	*course = "New Course"
	fmt.Println("\n Trying to change the course to", *course)
}