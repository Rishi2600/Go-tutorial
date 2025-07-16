package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var m sync.Mutex
var sum = 0

func main() {
	for i := range 10000 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer m.Unlock()
			m.Lock()

			sum += i
		}()
	}
	wg.Wait()

	fmt.Println("sum is: ", sum)
}
