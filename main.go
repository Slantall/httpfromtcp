package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}
	defer f.Close()

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")
	current_line := ""
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
			fmt.Printf("read: %s\n", split[i])
		}
		current_line = split[len(split)-1]
	}
	fmt.Printf("read: %s\n", current_line)
}
