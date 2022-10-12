package main

import (
	"fmt"
	"math/rand"
)

// Go channel
func Calculate(c chan int)  {
	c <- rand.Intn(10)
}
func main()  {
	fmt.Println("Go Channel Tutorial")
	values := make(chan int)
	defer close(values)

	go Calculate(values)
	value := <- values
	fmt.Println(value)
}
