package main

/**
 * The program displays the use of if statement and inline declaration in go
 *
 * @author Gaurav Gaur
 **/

import (
	"fmt"
)

func main() {

	if firstRank, secondrank := 34, 615; firstRank < secondrank {
		fmt.Println("first couse is popular" +
			"than second course")
	} else if firstRank > secondrank {
		fmt.Println("second course is popular than first course")
	} else {
		fmt.Println("both courses are same")
	}
	
}