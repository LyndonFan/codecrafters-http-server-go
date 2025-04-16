package main

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

func getEncoding(request *Request) (string, bool) {
	encoding, exists := request.Headers["Accept-Encoding"]
	if !exists {
		return "", false
	}
	if _, exists = ValidContentEncodings[encoding]; exists {
		return encoding, true
	} else {
		return "", false
	}
}