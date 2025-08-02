package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		sum := 0
		for i := range 55 {
			sum += i
		}
		ch <- sum
		defer wg.Done()
	}()
	go func() {
		sum := 0
		for i := 55; i < 101; i++ {
			sum += i
		}
		prev := <-ch
		fmt.Println(prev + sum)
		defer wg.Done()
	}()

	wg.Wait()

	fmt.Println("done")
}
