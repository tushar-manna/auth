//so I tried to implement http server from scratch

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listner, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	fmt.Println("We are currently listening on port 8000")

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			continue
		}
		go func(c net.Conn) {
			defer c.Close()

			//I will return just return a plain hello on https for any connection!
			response := "HTTP/1.1 200 OK\r\n" +
				"Content-Type: text/plain\r\n" +
				"Content-Length: 5\r\n" +
				"\r\n" +
				"Hello"

			c.Write([]byte(response))
		}(conn)
	}
}
