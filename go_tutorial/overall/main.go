package main

import (
	"fmt"
	"sync"
)

func expensiveOp() {
	sum := 0
	for i := range 1000 {
		sum += i
	}
	fmt.Printf("goroutine working, sum is: %v \n", sum)
}

func main() {
	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			expensiveOp()
		}()
	}

	wg.Wait()
	fmt.Println("main thread exists")
}
