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
	expensiveOp("first")
	expensiveOp("second")

	time.Sleep(time.Second * 2)

	fmt.Println("Done")
}
