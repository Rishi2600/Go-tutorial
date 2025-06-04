package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr2 := [3]int32{4, 4, 4} /*short hand to declare an array*/
	intArr3 := []int32{5, 6, 7}
	intArr[0] = 5

	fmt.Println(intArr2, intArr3)
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:2])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSliceArr []int32 = []int32{1, 2, 3}
	fmt.Printf("The length is %v with capactiy %v \n", len(intSliceArr), cap(intSliceArr))
	intSliceArr = append(intSliceArr, 556)
	fmt.Printf("The length is %v with capactiy %v \n", len(intSliceArr), cap(intSliceArr))
	fmt.Println(intSliceArr)

}

/*Array in Go has fixed length, Same type, indexable and contiguous memory.*/
/*in JavaScript, we can do "const mixedArray = [1, "string", {key: "value"}, function() {}]", but not in Go*/
/*in typescript, the approach is as same as Go, but can be managed accordingly "const arr: (string | number)[] = ['a', 1]"*/
/*Error: Argument of type '{ name: string; }' is not assignable to parameter of type 'string | number'*/
