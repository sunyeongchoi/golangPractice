package main

import (
	"context"
	"fmt"
	"time"
)

//// Adding values to our context
//func enrichContext(ctx context.Context) context.Context {
//	return context.WithValue(ctx, "api-key", "my-super-secret-api-key")
//}
//
//// Reading values from our context
//func doSomethingCool(ctx context.Context)  {
//	apiKey := ctx.Value("api-key")
//	fmt.Println(apiKey)
//}
//
//func main()  {
//	fmt.Println("Go Context Tutorial")
//	ctx := context.Background()
//	ctx = enrichContext(ctx)
//	doSomethingCool(ctx)
//}

// Context Deadlines using WithTimeout
func doSomethingCool(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("time out")
			err := ctx.Err()
			fmt.Println(err)
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Co Context Tutorial")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}
	time.Sleep(2 * time.Second)
}
