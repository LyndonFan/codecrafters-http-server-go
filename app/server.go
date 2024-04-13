package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const NEWLINE string = "\r\n"

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	argsWithoutProgram := os.Args[1:]
	directory := "" // leave empty if not supplied with "--directory" argument
	if len(argsWithoutProgram) == 2 && argsWithoutProgram[0] == "--directory" {
		directory = argsWithoutProgram[1]
	}
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
		go handleConnection(conn, directory)
	}
}

func handleConnection(conn net.Conn, directory string) error {
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return fmt.Errorf("expected a TCP connection, but failed to convert it")
	}
	input := make([]byte, 1024)
	length, err := conn.Read(input)
	if err != nil {
		return err
	}
	input = input[:length]
	request, err := parseRequest(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	response := handleRequest(request, directory)
	tcpConn.Write(response.Bytes())
	tcpConn.CloseWrite()
	return nil
}

func handleRequest(request *Request, directory string) Response {
	headers := map[string]string{}
	headers["Content-Type"] = "text/plain"
	response := Response{
		Version:       request.Version,
		StatusCode:    200,
		StatusMessage: "OK",
		Headers:       headers,
		Body:          []byte(""),
	}
	pathFields := strings.Split(request.Path[1:], "/") // omit first /
	switch {
	case request.Path == "/":
		// pass
	case request.Path == "/user-agent":
		response.Body = []byte(request.Headers["User-Agent"])
	case len(pathFields) == 2 && pathFields[0] == "echo":
		response.Body = []byte(pathFields[1])
	case len(pathFields) == 2 && pathFields[0] == "files":
		fullPath := fmt.Sprintf("%s/%s", directory, pathFields[1])
		content, err := os.ReadFile(fullPath)
		if os.IsExist(err) {
			response.Body = content
			response.Headers["Content-Type"] = "application/octet-stream"
		} else {
			response.StatusCode = 404
			response.StatusMessage = "Not Found"
		}
	default:
		response.StatusCode = 404
		response.StatusMessage = "Not Found"
	}
	headers["Content-Length"] = fmt.Sprintf("%d", len(response.Body))
	fmt.Println(len(response.Bytes()), string(response.Bytes()))
	return response
}
