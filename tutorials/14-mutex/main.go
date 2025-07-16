package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	m     sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	s := safeCounter{
		m:     sync.Mutex{},
		value: 0,
	}

	for i := range 100000 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer s.m.Unlock()
			s.m.Lock()
			s.value += i
		}()
	}
	wg.Wait()
	fmt.Print("sum is ", s.value)
}
