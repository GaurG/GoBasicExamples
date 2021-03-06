package main

/**
 * The program displays the basic usage of gorountines.
 * Goroutines help us achieve concurrency.
 *
 * @author Gaurav Gaur
 * 
 **/

 import (
 	"fmt"
 	"time"
 	"sync"
 )

 func main() {

 	// This will allow main to wait for other 2 routines to complete
 	var waitGrp sync.WaitGroup
 	waitGrp.Add(2)

 	/*
 	 * Couple things to notice in this function:
 	 * 1. go keyword marks the function for concurrent execution. go keyword
 	 * will trigger a routine
 	 * 2. () braces at the end of the function marks the function as anonymous
 	 *     and the function will be executed when main executes.
 	 * 3. Ideally Hello should be printed before my name.
	 */
 	go func() {
 		defer waitGrp.Done()

 		time.Sleep(5 * time.Second)
 		fmt.Println("Hello")
 	}()

 	go func() {
 		defer waitGrp.Done()
 		fmt.Println("Gaurav Gaur")
 	}()

 	waitGrp.Wait()
 }