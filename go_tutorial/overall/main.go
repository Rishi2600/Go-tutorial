package main

import (
	"fmt"
)

func main() {

	var ch = make(chan int)

	go func() {
		sum := 0
		for i := range 50 {
			sum += i
		}
		ch <- sum
	}()
	firstSum := <-ch
	fmt.Println(firstSum)

	go func() {
		sum := 0
		for i := 50; i < 101; i++ {
			sum += i
		}
		ch <- sum
	}()
	secondSum := <-ch
	fmt.Println(secondSum)

	fmt.Println(firstSum + secondSum)

}
