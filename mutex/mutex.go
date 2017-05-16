package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var mut = sync.RWMutex{}

var counter = 0 // this will be shared by multiple go routines , so mutex is needed

func main() {
	// runtime.GOMAXPROCS(100)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go increment()
		go print()
	}
	wg.Wait()
}

func increment() {
	counter++
	wg.Done()
}

func print() {
	fmt.Printf("#%v\n", counter)
	wg.Done()
}
