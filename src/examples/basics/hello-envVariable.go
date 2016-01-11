package main

/**
 * The program uses env variable
 *
 * @author Gaurav Gaur
 *
 **/
import (
	"fmt"
	"os" // package to access env varaibles
)

func main() {
	
	name := os.Getenv("USERNAME")
	course := "Go fundamentals"

	fmt.Println("\nHi", name, "you are currently watching", course)
	changeCourse(&course)
	fmt.Println("\nYou are now watching the course", course)
	
}

func changeCourse(course *string) {
	*course = "New Course"
	fmt.Println("\n Trying to change the course to", *course)
}