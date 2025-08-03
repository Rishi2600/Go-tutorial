package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func expensiveOp(wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := range 10 {
		sum += i
	}
	fmt.Printf("goroutine working, sum is: %v \n", sum)
}

func main() {

	wg.Add(1)
	go expensiveOp(&wg)

	wg.Wait()
	fmt.Println("main thread exit")
}
