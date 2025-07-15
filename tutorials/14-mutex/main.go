package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			counter++
		}()

	}

	wg.Wait()

	fmt.Println("final counter: ", counter)

}
