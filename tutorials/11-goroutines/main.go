package main

import (
	"fmt"
	"time"
)

func expensiveOp(str string) {
	for i := range 3 {
		fmt.Println(str, "-", i)
	}
}

func main() {

	time1 := time.Now()

	go expensiveOp("first")
	go expensiveOp("second")
	go expensiveOp("third")
	go expensiveOp("fourth")
	go expensiveOp("fifth")
	go expensiveOp("sixth")
	go expensiveOp("seventh")
	go expensiveOp("eighth")

	time.Sleep(time.Second * 2)

	fmt.Println("Done")

	timeTaken := time.Since(time1)

	fmt.Printf("time taken: %v \n", timeTaken)
}
