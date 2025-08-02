package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		sum := 0
		for i := range 55 {
			sum += i
		}
		ch <- sum
	}()
	go func() {
		sum := 0
		for i := 55; i < 101; i++ {
			sum += i
		}
		prev := <-ch
		fmt.Println(prev + sum)
	}()

	time.Sleep(time.Second * 2)

	fmt.Println("done")
}
