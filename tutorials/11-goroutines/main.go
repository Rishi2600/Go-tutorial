package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 100 {
			fmt.Println("return from the first function:", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range 50 {
			fmt.Println("return from the second function:", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}()

	wg.Wait()
	fmt.Println("Done! All the funcions called")
}
