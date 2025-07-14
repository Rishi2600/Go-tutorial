package main

import (
	"fmt"
	"sync"
	"time"
)

func expOperation() {
	sum := 0
	for i := range 100000000 {
		sum += i
	}
	println(sum)
}

func main() {

	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			expOperation()
		}()
	}

	time.Sleep(time.Millisecond * 2000)

	wg.Wait()
	fmt.Println("concurrent processes done")

	time.Sleep(time.Millisecond * 5000)

	fmt.Println("Done! main func called")
}
