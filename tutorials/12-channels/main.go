package main

import (
	"fmt"
	"time"
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

	go sender(c)
	go receiver(c)

	time.Sleep(time.Millisecond * 2000)

}
