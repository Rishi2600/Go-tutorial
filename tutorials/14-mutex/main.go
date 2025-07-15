package main

import (
	"fmt"
	"sync"
)

var counter int
var m sync.Mutex

func main() {
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			m.Lock()
			counter++
			m.Unlock()
		}()

	}

	wg.Wait()

	fmt.Println("final counter: ", counter)

}
