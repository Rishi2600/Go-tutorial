package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array [3]string = [3]string{"1", "2", "3"}
	fmt.Println(reflect.TypeOf(array))
	fmt.Println(array)

	var slice = array[:]
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(slice)

}
