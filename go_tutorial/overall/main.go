package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		for i := range 10 {
			fmt.Printf("iterated value: %v \n", i)
			ch <- i
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Printf("channel value: %v \n", <-ch)
}
