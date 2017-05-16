package main

import (
	"fmt"
)

type Books struct {
	title  string
	author string
}

//reciever type
func (b *Books) p() {
	fmt.Printf("Book title : %s\n", b.title)
}

//value type
func (b Books) p2() {
	fmt.Printf("Book title : %s\n", b.title)
}

func main() {

	var Book1 Books
	var Book2 Books

	Book1.title = "Go Programming"

	Book2.title = "Telecom Billing"

	Book1.p()
	Book1.p2()
}

func printBookP(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
}
