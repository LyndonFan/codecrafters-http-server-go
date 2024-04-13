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
	defer l.Close()
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
		return fmt.Errorf("expected a TCP connection, but failed to convert it")
	}
	request, err := parseRequest(tcpConn)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	headers := map[string]string{}
	headers["Content-Type"] = "text/plain"
	response := Response{
		Version:       request.Version,
		StatusCode:    200,
		StatusMessage: "OK",
		Headers:       headers,
		Body:          []byte(""),
	}
	if len(request.Path) >= 6 && request.Path[:6] == "/echo/" {
		response.Body = []byte(request.Path[6:])
	} else if request.Path == "/user-agent" {
		response.Body = []byte(request.Headers["User-Agent"])
	} else if request.Path != "/" {
		response.StatusCode = 404
		response.StatusMessage = "Not Found"
	}
	headers["Content-Length"] = fmt.Sprintf("%d", len(response.Body))
	fmt.Println(len(response.Bytes()), string(response.Bytes()))
	tcpConn.Write(response.Bytes())
	tcpConn.CloseWrite()
	return nil
}
