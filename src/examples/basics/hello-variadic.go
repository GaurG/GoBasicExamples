package main

/**
 * The program demonstrates the use of multiple parameter(...)
 *
 * @author Gaurav Gaur
 *
 **/

import (
	"fmt"
)

func main() {

	bestFinish := bestLeagueFinish(10,11,15,17,14,16)
	fmt.Println(bestFinish)
}

func bestLeagueFinish(finishes ...int) int {

	best := finishes[0]
	for _,i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}