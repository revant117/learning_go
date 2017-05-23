package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) //channel is unbuffered by default
	// unbuffered: channel will only send if there is something to recieve , else deadlock error
	wg.Add(2)
	go pushinto(ch)
	go pullfrom(ch)
	wg.Wait()

}

func pushinto(ch chan<- int) { // can use (ch chan int ) also but this way we can see if which end of the chanel we are using
	i := 42
	ch <- i
	wg.Done()
}

func pullfrom(ch <-chan int) {
	a := <-ch
	fmt.Println(a)
	wg.Done()
}
