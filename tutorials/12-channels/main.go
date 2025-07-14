package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		sum := 0
		for i := range 100 {
			fmt.Println("return from the first function: ", i)
			sum += i
		}
		c <- sum
	}()

	output := <-c

	fmt.Println("Output: ", output)
}
