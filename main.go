package main

//get conn
//accept conn
//return status

import (
	"fmt"
	"net"
)

func ListeningAndServing(ln net.Listener, port int) {
	var err error
	ln, err := net.Listen("tcp", ":6060")
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	fmt.Println("i am listening.....")
}
