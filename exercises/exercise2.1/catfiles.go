package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file_names := os.Args[1:]

	for _, file_name := range file_names {
		go func(fname string) {
			file_content, err := os.ReadFile(fname)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s", file_content)
		}(file_name)
	}

	time.Sleep(3 * time.Second)
}
