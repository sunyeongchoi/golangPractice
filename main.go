package main

import (
	"fmt"
	"time"
)

// Goroutine
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func main() {
	fmt.Println("Goroutine Tutorial")
	// notice how we've added the 'go' keyword
	// in front of both our compute function calls
	go compute(10)
	go compute(10)

	go func() {
		fmt.Println("Executing my Concurrent anonymous function")
	}()

	// we scan fmt for input and print that to our console
	// main function completed before our asynchronous functions could execute and as such, any goroutines that have yet to complete are promptly terminated.
	// our program waits for keyboard input before it kills off our poor goroutines
	var input string
	fmt.Scanln(&input)
}
