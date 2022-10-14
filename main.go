package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroups

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Inside my goroutine")
	// signal the end of its' execution
	wg.Done()
}

func main() {
	fmt.Println("Hello World")
	var waitgroup sync.WaitGroup
	// set the number of goroutines
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	// block the execution of our main() function until the goroutines in the waitgroup have successfully completed
	waitgroup.Wait()
	fmt.Println("Finish Execution")
}
