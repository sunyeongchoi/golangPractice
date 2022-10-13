package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go channel - Buffered Channels
func Calculate(c chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("This executes regardless as the send is now non-blocking")
}
func main() {
	fmt.Println("Go Channel Tutorial")
	valueChannel := make(chan int, 2)
	defer close(valueChannel)

	go Calculate(valueChannel)
	go Calculate(valueChannel)

	values := <-valueChannel
	fmt.Println(values)
}
