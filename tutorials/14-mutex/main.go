package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var (
	data = make(map[int]int)
	m    sync.Mutex
)

func writeToMap(i int) {
	defer wg.Done()
	defer m.Unlock()
	m.Lock()
	data[i] = i + 1
}

func main() {
	for i := range 1000 {
		wg.Add(1)

		go writeToMap(i)
	}

	wg.Wait()

	fmt.Println("Done writing to map, length of map:", len(data))
	fmt.Println("data:", data)
}
