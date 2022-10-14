package main

import (
	"fmt"
)

// panic

func sillySusan() {
	fmt.Println("silly susan called")
	panicKingPeter()
	fmt.Println("silly susan finished")
}

func panicKingPeter() {
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("panicking peter called")
	// It checks to see if there is anything deferred within panickingPeter and executes these deferred statements.
	// It then terminates.
	panic("oh no")
	fmt.Println("panicking peter finished")
}
func main() {
	fmt.Println("cascading panics")
	sillySusan()
}
