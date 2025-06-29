package main

import "fmt"

func main() {
	var a int = 5
	fmt.Println(a)

	var b int = a
	fmt.Println(b)
	var bPtr *int = &b
	fmt.Println(bPtr)

	var z *int = &a
	fmt.Println(z)
	fmt.Println(*z)

	*z = 9
	fmt.Println(*z)
	fmt.Println(a)
	fmt.Println(b)

}

/*& defines the address and * goes for the pointer, a value of a var can be cahnged by just cahnging the value of the underlying address*/
