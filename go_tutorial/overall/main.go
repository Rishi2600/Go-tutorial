package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func expensiveOp(ch chan int) {
	sum := 0
	for i := range 10 {
		sum += i
	}

	ch <- sum
}

func main() {

	go expensiveOp(ch)

	finalSum := <-ch

	fmt.Println(finalSum)

	time.Sleep(time.Millisecond * 2000)
	fmt.Println("done")
}
