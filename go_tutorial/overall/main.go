package main

import "fmt"

func main() {
	var object = make(map[string]string)

	object["age1"] = "56"
	object["age2"] = "6"
	object["age3"] = "5"
	object["age4"] = "567"

	var _, exists = object["map3"]
	var value, existss = object["age4"]

	println(object)
	fmt.Println(object, exists)
	fmt.Printf("exists: %v & value: %v", existss, value)

}
