package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("Error creating listener: %w\n", err)
	}
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error with listener: %w\n", err)
		}
		//fmt.Printf("Connection accepted: %v\n", connection)
		str_ch := getLinesChannel(connection)
		for s := range str_ch {
			fmt.Printf("read: %s\n", s)
		}
		_, ok := <-str_ch
		if !ok {
			fmt.Println("Connection has been closed")
		}

	}

}
