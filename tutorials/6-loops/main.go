package main

import "fmt"

func main() {
	var value = 0
	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		value += i
	}
	fmt.Println(value)
}

/*i:= range 10 means i := 0; i < 10; i++*/
