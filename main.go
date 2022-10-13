package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go channel - Unbuffered Channels
func Calculate(c chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}
func main() {
	fmt.Println("Go Channel Tutorial")
	values := make(chan int)
	defer close(values)

	go Calculate(values)
	go Calculate(values)

	value := <-values
	fmt.Println(value)
}
