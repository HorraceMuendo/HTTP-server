package connectionhandler

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

	fmt.Println("listening on port: %s host:%s \n", port, host)

}
