package main

import (
	"fmt"
	"time"
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
	time.Sleep(time.Second * 5)
	firstSum := <-ch

	go func() {
		sum := 0
		for i := 50; i < 101; i++ {
			sum += i
		}
		ch <- sum
	}()
	secondSum := <-ch

	fmt.Println(firstSum, secondSum, firstSum+secondSum)

}
