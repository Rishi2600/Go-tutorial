package main

import "fmt"

func main() {
	var myString = []rune("rèsumè")
	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Println(indexed)
}

/*rune is a data type that represents a Unicode code point*/
/*It stores a 32-bit integer value that corresponds to a specific character in the Unicode standard*/
