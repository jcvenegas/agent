package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("missing args <ip> <cmd>")
		os.Exit(1)
	}
	ip := os.Args[1]
	cmd := os.Args[2]

	conn, err := net.Dial("tcp", ip+":8081")
	if err != nil {
		fmt.Println("failed to connect ", err)
		os.Exit(1)
	}
	defer conn.Close()

	// send to socket
	_, err = fmt.Fprintf(conn, cmd+"\n")
	if err != nil {
		fmt.Println("failed to write to conn ", err)
		os.Exit(1)
	}
	tmp := make([]byte, 256) // using small tmo buffer for demonstrating
	for {
		_, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)

			}
			break

		}
		fmt.Print(string(tmp))
	}
}
