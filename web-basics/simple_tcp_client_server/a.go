package main

import (
	"fmt"
	"net"
	"log"
	"io"
)

func main() {
	li , err := net.Listen("tcp" , ":8080")
	if err != nil{
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn , err  :=  li.Accept()
		if err != nil{
			log.Panic(err)
		}
		fmt.Fprintln(conn , "connection")
		io.WriteString(conn , "\nHello from tcp\n")
		conn.Close()
	}
}	