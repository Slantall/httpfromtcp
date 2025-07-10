package main

import (
	"fmt"
	"log"
	"os"
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
	str_ch := getLinesChannel(f)
	for s := range str_ch {
		fmt.Printf("read: %s\n", s)
	}

}
