package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var data = make(map[int]int)

func writeToMap(i int) {
	defer wg.Done()
	data[i] = i
}

func main() {
	for i := range 10000 {
		wg.Add(1)

		go writeToMap(i)
	}

	wg.Wait()

	fmt.Println("Done writing to map")
}
