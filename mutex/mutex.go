package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var mut = sync.RWMutex{} // Read Write Mutex , it locks data so that only one method can use it at a time

var counter = 0 // this will be shared by multiple go routines , so mutex (lock) is needed
//otherwise racing will occure between print and increment methods

func main() {
	runtime.GOMAXPROCS(100)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		mut.Lock() // lock "writing data" here  , prevents racing
		go increment()
		mut.RLock()
		go print()
	}
	wg.Wait()

	// IMP: but by adding mutex , concurency has been killed , so this example is just to understand how mutex works
}

func increment() {
	counter++
	mut.Unlock() // Unlock writing data here
	wg.Done()
}

func print() {
	fmt.Printf("#%v\n", counter)
	mut.RUnlock()
	wg.Done()
}
