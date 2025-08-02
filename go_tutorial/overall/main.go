package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := range 1000 {
			fmt.Println(i)
		}
	}()
	go func() {
		for i := range 1000 {
			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second * 2)

	fmt.Println("done")
}
