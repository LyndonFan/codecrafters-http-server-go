package main

import "testing"

func assertRequestsEqual(t *testing.T, expected, actual Request) {
	if actual.Method != expected.Method {
		t.Fail()
		t.Logf("Expected Method %s, got Method %s", expected.Method, actual.Method)
	}
	if actual.Path != expected.Path {
		t.Fail()
		t.Logf("Expected Path %s, got Path %s", expected.Path, actual.Path)
	}
	if actual.Version != expected.Version {
		t.Fail()
		t.Logf("Expected Version %s, got Version %s", expected.Version, actual.Version)
	}
	if actual.Body != expected.Body {
		t.Fail()
		t.Logf("Expected Body %s, got Body %s", expected.Body, actual.Body)
	}
}

func TestRequestNoHeaders(t *testing.T) {
	requestString := "GET / HTTP/1.1\r\n\r\n\r\n"
	expected := Request{
		Method:  "GET",
		Path:    "/",
		Version: "HTTP/1.1",
		Headers: map[string]string{},
		Body:    "",
	}
	actual, err := parseRequest([]byte(requestString))
	if err != nil {
		t.Fatalf("Expected no errors, got %v", err)
	}
	assertRequestsEqual(t, expected, *actual)
}
