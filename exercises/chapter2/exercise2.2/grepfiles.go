package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func searchFile(search_word, fname string) {
	file_content, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Fields(string(file_content))
	for _, word := range words {
		if search_word == word {
			fmt.Printf("%s contains a match", fname)
			break
		}
	}
}

func main() {
	file_names := os.Args[2:]
	for _, file_name := range file_names {
		go searchFile(os.Args[1], file_name)
	}

	time.Sleep(3 * time.Second)
}
