package main

import (
	"fmt"
)

func sender(c chan string) {
	c <- "this is the message from the sender"
}

func receiver(c chan string) {
	message := <-c
	fmt.Printf("message received: %v", message)
}

func main() {

	c := make(chan string)

	defer close(c)

	go sender(c)

	go receiver(c)
}
