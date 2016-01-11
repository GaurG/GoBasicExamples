package main

/**
 * The program proves that GO uses pass by value and not
 * pass by reference for primitive types
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
	changeCourse(course)
	fmt.Println("\nYou are now watching the course", course)
	
}

func changeCourse(course string) string {
	course = "New Course"
	fmt.Println("\n Trying to change the course to", course)
	return course
}