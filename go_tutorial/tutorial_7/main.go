package main

import "fmt"

func main() {
	var userDetails *int32 = new(int32)
	fmt.Println(userDetails)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4

	fmt.Println(slice)
	fmt.Println(sliceCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing1 array is: %p \n", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("The result is: %v \n", result)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing1 array is: %p \n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

/*pointer is a variable that stores the memory address of another variable*/
/*pointer allows direct manipulation of memory and provides a way to access and modify the value at the referenced address*/
/*if you have an integer variable x, you can declare a pointer to it using *int,
	and assign the memory address of x to the pointer using the & operator. eg -
var x int = 10
var p *int
p = &x*/
