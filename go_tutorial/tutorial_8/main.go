package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}
	fmt.Printf("total execution time: %v \n", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("the responde from the DB is: ", dbData[i])
}

/*with using GoRoutine, we can achieve some level of parallel execution as long as there is a mutli core CPU(we usually do have it).*/
