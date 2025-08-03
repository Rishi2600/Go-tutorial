package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool)

	go func() {
		sum := 0
		for i := range 10 {
			sum += i
		}
		fmt.Printf("goroutine working, sum is: %v \n", sum)
		ch <- true
	}()

	<-ch
	fmt.Println("main thread exit")
}
