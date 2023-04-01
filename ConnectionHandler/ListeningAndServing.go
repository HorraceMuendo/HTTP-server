package connectionhandler

// starts a server that listens on port 5000 for incoming connections

import (
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
