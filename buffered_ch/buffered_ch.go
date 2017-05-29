package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	//channel is unbuffered by default , but if we add second arg then it becomes buffered , here it can take 50 integers
	// unbuffered: channel will only send if there is something to recieve , else deadlock error
	wg.Add(2)
	go pushinto(ch)
	go pullfrom(ch)
	wg.Wait()

}

// send only
// best practise to use only one type of channel in one go routine , avoid using both sending and recieving form
//the same channel in the same go routine
func pushinto(ch chan<- int) { // can use (ch chan int ) also but this way we can see if which end of the chanel we are using
	i := 42
	ch <- i
	ch <- 35
	close(ch) // close the channel , for example in case we need to loop over this channel we need to use close() else
	// it will cause a dead lock in the recieveing go routine , see line no. 42
	//now this would causes a dead lock in case of unbuffered channel
	wg.Done()
}

//recieve only
func pullfrom(ch <-chan int) {
	a := <-ch
	fmt.Println(a)
	a = <-ch // this gets 35 , the second value sent into the channel
	fmt.Println(a)

	// same thinng we did here can be done by looping over a channel , for example :-
	// for i := range ch {
	// 	fmt.Println(i)
	// }
	wg.Done()
}
