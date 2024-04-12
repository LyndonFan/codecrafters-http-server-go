package main

import (
	"fmt"
	"net"
	"os"
)

const NEWLINE string = "\r\n"

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
	request, err := parseRequest(conn)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	headers := map[string]string{}
	headers["Content-Type"] = "text/plain"
	var returnString string
	if len(request.Path) >= 6 && request.Path[:6] == "/echo/" {
		returnString = request.Path[6:]
	}
	headers["Content-Length"] = fmt.Sprintf("%d", len(returnString))
	response := Response{
		Version:       request.Version,
		StatusCode:    200,
		StatusMessage: "OK",
		Headers:       headers,
		Body:          []byte(returnString),
	}
	conn.Write(response.Bytes())
	err = conn.Close()
	if err != nil {
		fmt.Printf("Error when closing: %v\n", err)
	} else {
		fmt.Println("Closed successfully")
	}
	return nil
}
