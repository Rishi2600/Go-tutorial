package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func expensiveOp(str string) {

	defer wg.Done()

	for i := range 3 {
		fmt.Println(str, "-", i)
	}
}

func main() {

	time1 := time.Now()

	wg.Add(2)

	go expensiveOp("first")
	go expensiveOp("second")

	wg.Wait()

	fmt.Println("Done")

	timeTaken := time.Since(time1)
	fmt.Printf("time taken: %v \n", timeTaken)
}
