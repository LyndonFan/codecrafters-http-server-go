package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Response struct {
	Version       string
	StatusCode    int
	StatusMessage string
	Headers       map[string]string
	Body          []byte
}

func (resp *Response) Bytes() []byte {
	lines := make([]string, 1, 3)
	lines[0] = strings.Join(
		[]string{resp.Version, strconv.Itoa(resp.StatusCode), resp.StatusMessage},
		" ",
	)
	for k, v := range resp.Headers {
		lines = append(lines, fmt.Sprintf("%s: %s", k, v))
	}
	if len(resp.Body) > 0 {
		lines = append(lines, "", string(resp.Body))
	}
	return []byte(strings.Join(lines, NEWLINE))
}
