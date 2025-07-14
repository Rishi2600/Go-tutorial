package main

import (
	"fmt"
)

func expOperation(ch chan<- bool) {
	sum := 0
	for i := range 100000000 {
		sum += i
	}
	println(sum)
	ch <- true
}

func main() {

	ch1 := make(chan bool)
	go expOperation(ch1)

	ch2 := make(chan bool)
	go expOperation(ch2)

	ch3 := make(chan bool)
	go expOperation(ch3)

	ch4 := make(chan bool)
	go expOperation(ch4)

	ch5 := make(chan bool)
	go expOperation(ch5)

	<-ch1
	<-ch2
	<-ch3
	<-ch4
	<-ch5

	fmt.Println("Done")
}
