package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	str_ch := make(chan string)
	current_line := ""
	go func(f io.ReadCloser) <-chan string {
		defer f.Close()
		defer close(str_ch)
		for {
			b := make([]byte, 8)
			n, err := f.Read(b)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}
			str := string(b[:n])
			current_line += str
			split := strings.Split(current_line, "\n")
			for i := 0; len(split)-i > 1; i++ {
				str_ch <- split[i]
			}
			current_line = split[len(split)-1]
		}
		str_ch <- current_line

		return str_ch
	}(f)

	return str_ch
}
