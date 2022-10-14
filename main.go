package main

import (
	"fmt"
	"sync"
)

// WaitGroups

func main() {
	fmt.Println("Hello World")
	var waitgroup sync.WaitGroup
	// set the number of goroutines
	waitgroup.Add(1)
	go func() {
		fmt.Println("Inside my goroutine")
		// signal the end of its' execution
		waitgroup.Done()
	}()
	// block the execution of our main() function until the goroutines in the waitgroup have successfully completed
	waitgroup.Wait()
	fmt.Println("Finish Execution")
}
