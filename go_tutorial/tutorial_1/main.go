package main

import "fmt"

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var intNum1 int8 = 56
	var intNum2 int16 = 66

	var finalInt1 = intNum1 + int8(intNum2)
	fmt.Println(finalInt1)

	var floatNum float64 = 56.5486589
	fmt.Println(floatNum)

	var intNum3 int = 522
	var floatNum3 float64 = 56.451

	var finalInt2 = float64(intNum3) + floatNum3
	fmt.Println(finalInt2)

	var myString string = "this is a string"
	fmt.Println(myString)

	var myString1 = "string1" /*type is infered here*/
	fmt.Println(myString1)

	myString2, myStirng3 := "string2", "string3" /*short hand for defining var or const*/
	fmt.Println(myString2, myStirng3)

	/*var vs const*/
	var variable string
	const constant string = "const needs an init value but var doesn't"
	fmt.Println(variable, constant)
}
