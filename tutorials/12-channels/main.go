package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5)

	go func() {
		for i := range 10 {
			fmt.Println("return from the first function: ", i)
			c <- i
		}
	}()

	time.Sleep(time.Millisecond * 1000)

	for values := range c {
		fmt.Println(values)
	}

}
