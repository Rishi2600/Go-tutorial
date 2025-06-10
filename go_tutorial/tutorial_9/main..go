package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_PRICE float32 = 5

func main() {
	var myChannel = make(chan string)
	var websites = []string{"amazon", "flipkart", "walmart"}

	for i := range websites {
		go checkPrice(websites[i], myChannel)
	}
	sendMessage(myChannel)
}

func checkPrice(websites string, myChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price <= MAX_PRICE {
			myChannel <- websites
			break
		}
	}
}

func sendMessage(myChannel chan string) {
	fmt.Printf("found a deal on %v \n", <-myChannel)
}

/*channel helps to make the process parallel working along with goroutines*/
