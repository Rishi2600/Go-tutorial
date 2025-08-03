package main

import (
	"fmt"
)

func main() {

	go func() {
		sum := 0
		for i := range 10 {
			sum += i
		}
		fmt.Printf("goroutine working, sum is: %v \n", sum)
	}()

	fmt.Println("main thread exit")
}
