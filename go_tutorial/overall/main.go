package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sum := 0

	for range 1000 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for range 100 {
				sum += 1
			}
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}
