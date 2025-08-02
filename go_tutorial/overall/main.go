package main

import (
	"fmt"
	"time"
)

func expensiveOp(ch chan int) {
	sum := 0
	for i := range 10 {
		sum += i
	}

	ch <- sum
}

func main() {

	ch := make(chan int)

	go expensiveOp(ch)

	finalSum := <-ch

	fmt.Println(finalSum)

	time.Sleep(time.Millisecond * 2000)
	fmt.Println("done")
}
