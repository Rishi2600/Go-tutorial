package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5)

	go func() {
		sum := 0
		for i := range 10 {
			fmt.Println("return from the first function: ", i)
			sum += i
		}
		c <- sum
	}()

	time.Sleep(time.Millisecond * 1000)
	output := <-c

	fmt.Println("Output: ", output)
}
