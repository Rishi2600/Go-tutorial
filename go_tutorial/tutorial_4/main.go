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

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]string{"name": "rishi", "age": "23"} /*short hand for map --key value pair.*/
	fmt.Println(myMap2)
	fmt.Println(myMap2["name"])
	delete(myMap, "age")

	for name := range myMap2 {
		fmt.Printf("Name: %v \n", name)
	}

	var i int
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

