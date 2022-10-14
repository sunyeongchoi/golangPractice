package main

import (
	"fmt"
	"time"
)

// WaitGroups

func myFunc() {
	time.Sleep(1 * time.Second)
	fmt.Println("Inside my goroutine")
}

// myFunc 고루틴 함수가 실행될 기회를 주지 않고 main 함수가 종료된다.
func main() {
	fmt.Println("Hello World")
	go myFunc()
	fmt.Println("Finish Execution")
}
