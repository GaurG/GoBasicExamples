package main

/**
 * The program displays basic for loop
 *
 * @author Gaurav Gaur
 *
 **/

import (
	"fmt"
	"time"
)

func main() {

	for counter := 10; counter >= 0; counter-- {
		fmt.Println(counter)
		time.Sleep(1 * time.Second)
	}
}