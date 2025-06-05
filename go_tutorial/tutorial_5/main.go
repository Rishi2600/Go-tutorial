package main

import "fmt"

func main() {
	var myString = []rune("rèsumè")
	var indexed = myString[1]
	fmt.Println(myString)
	fmt.Println(indexed)
}
