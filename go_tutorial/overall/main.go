package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sum := 0
		for i := range 10 {
			sum += i
		}
		fmt.Printf("goroutine working, sum is: %v \n", sum)
	}()

	wg.Wait()
	fmt.Println("main thread exit")
}

/*it's the most ideal way of making sure the goroutine is completed before the main thread exists*/
