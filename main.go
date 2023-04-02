package main

import (
	//	connectionhandler "connectionhandler/ListeningAndServing"
	"fmt"
	"net"
)

func ListeningAndServing() {
	var err error
	port := "localhost:5000"

	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()

	host, port, err := net.SplitHostPort(ln.Addr().String())

	fmt.Println("listening on port: %s host: %s \n", port, host)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go func(conn net.Conn) {
			buffer := make([]byte, 1024)
			len, err := conn.Read(buffer)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Message recieved: %s\n ", string(buffer[:len]))
			conn.Write([]byte("message recieved .\n"))
			conn.Close()
		}(conn)
	}

}

func AcceptDial() {
	address := ""

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Unable to establish a connection", err)
	}
	fmt.Println("connection established on address: %s \n", conn)
}

func main() {
	//connectionhandler.ListeningAndServing()
	ListeningAndServing()
	AcceptDial()
	fmt.Println("i am listening.....")
}
