package main

import "fmt"

func main() {
	var value = 0
	for i := range 10 {
		if i%2 == 0 {
			break
		}
		value += i
		fmt.Println(value)
	}
	fmt.Println("Hello, outside of the loop")
}

/*i:= range 10 means i := 0; i < 10; i++*/
