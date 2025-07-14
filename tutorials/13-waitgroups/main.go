package main

import (
	"fmt"
	"sync"
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

	wg.Wait()

	fmt.Println("Done")
}
