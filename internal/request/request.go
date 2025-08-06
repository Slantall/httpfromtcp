package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine `json:"requestLine"`
}

type RequestLine struct {
	HttpVersion   string `json:"httpVersion"`
	RequestTarget string `json:"requestTarget"`
	Method        string `json:"method"`
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	data, err := io.ReadAll(reader) // Simulating reading from the reader
	if err != nil {
		return nil, fmt.Errorf("failed to read request: %w", err)
	}

	lines := strings.Split(string(data), "\r\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("no request line found")
	}

	requestLine := lines[0]
	parsedLine, err := parseRequestLine(requestLine)
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: parsedLine,
	}, nil

}

func parseRequestLine(line string) (RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return RequestLine{}, fmt.Errorf("invalid request line: %s", line)
	}
	method := parts[0]
	if method != strings.ToUpper(method) {
		return RequestLine{}, fmt.Errorf("unexpected method: %s", method)
	}
	httpVersion := strings.Split(parts[2], "/")[1]
	if httpVersion != "1.1" {
		return RequestLine{}, fmt.Errorf("unsupported HTTP version: %s", httpVersion)
	}
	return RequestLine{
		Method:        method,
		RequestTarget: parts[1],
		HttpVersion:   strings.Split(parts[2], "/")[1],
	}, nil
}
