package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		sum := 0
		for i := range 10 {
			sum += i
		}
		fmt.Printf("goroutine working, sum is: %v \n", sum)
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("main thread exit")
}
