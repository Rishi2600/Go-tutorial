package main

import (
	"fmt"
	"time"
)

func main() {

	sum := 0

	go func() {
		for i := range 3 {
			sum += i
			fmt.Println(sum)
		}
		fmt.Println(sum)
	}()
	go func() {
		for i := range 3 {
			sum += i
			fmt.Println(sum)
		}
		fmt.Println(sum)
	}()

	time.Sleep(time.Second * 2)

	fmt.Println("done")
}
