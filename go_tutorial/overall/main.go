package main

import (
	"fmt"
)

func main() {
	func() {
		for i := range 1000 {
			fmt.Println(i)
		}
	}()
	func() {
		for i := range 1000 {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}
