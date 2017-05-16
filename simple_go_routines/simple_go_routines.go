package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

//waitGroup synchronozes between 2 go routines

func main() {

	msg := "hey"
	wg.Add(1)

	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
		// Tell waitgroup when this go routine is done
	}(msg)
	msg = "goodbye"

	wg.Wait()
}
