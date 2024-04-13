package main

import (
	"fmt"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

func parseRequest(input []byte) (*Request, error) {
	fmt.Println(len(input), string(input))
	lines := strings.Split(string(input), NEWLINE)
	if len(lines) < 3 {
		return nil, fmt.Errorf("invalid request")
	}
	firstLine := strings.Split(lines[0], " ")
	if len(firstLine) != 3 {
		return nil, fmt.Errorf("invalid request")
	}
	method := firstLine[0]
	path := firstLine[1]
	version := firstLine[2]
	headers := make(map[string]string)
	for i := 1; i < len(lines)-2; i++ {
		if lines[i] == "" {
			continue
		}
		line := strings.Split(lines[i], ": ")
		if len(line) < 2 {
			return nil, fmt.Errorf("invalid request, line %s isn't a valid header", line)
		}
		headers[line[0]] = strings.Join(line[1:], ": ")
	}
	if lines[len(lines)-2] != "" {
		return nil, fmt.Errorf("invalid request, second last line is nonempty")
	}
	body := lines[len(lines)-1]
	request := Request{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: headers,
		Body:    body,
	}
	return &request, nil
}
