package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			break
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) error {
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return fmt.Errorf("connection is not a TCP connection")
	}
	request, err := parseRequest(tcpConn)
	if err != nil {
		return err
	}
	var response string
	if request.Method == "GET" && request.Path == "/" {
		response = "HTTP/1.1 200 OK\r\n\r\n"
	} else {
		response = "HTTP/1.1 404 Not Found\r\n\r\n"
	}
	tcpConn.Write([]byte(response))
	tcpConn.CloseWrite()
	return nil
}
