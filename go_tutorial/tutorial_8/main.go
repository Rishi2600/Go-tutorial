package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var waitGroups = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := range dbData {
		waitGroups.Add(1)
		go dbCall(i)
	}
	waitGroups.Wait()
	fmt.Printf("total execution time: %v \n", time.Since(t0))
	fmt.Printf("the results are: %v \n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("the responde from the DB is: ", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	waitGroups.Done()
}

/*with using GoRoutine, we can achieve some level of parallel execution.*/
