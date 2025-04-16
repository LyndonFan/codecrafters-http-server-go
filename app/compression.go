package main

import "strings"

var ValidContentEncodings = map[string]bool{
	"br": true,
	"compress": true,
	"deflate": true,
	"exi": true,
	"gzip": true,
	"identity": true,
	"pack200-gzip": true,
	"zstd": true,
}

func getEncodings(request *Request) []string {
	encodingString, exists := request.Headers["Accept-Encoding"]
	if !exists {
		return []string{}
	}
	possibleEncodings := strings.Split(encodingString, ",")
	res := make([]string, 0, len(possibleEncodings))
	for _, pe := range possibleEncodings {
		pe = strings.TrimSpace(pe)
		if _, exists = ValidContentEncodings[pe]; exists {
			res = append(res, pe)
		}
	}
	return res
}