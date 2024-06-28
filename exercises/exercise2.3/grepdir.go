package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func grepPath(path string, fileInfo fs.DirEntry, searchStr string) {
	fullPath := filepath.Join(path, fileInfo.Name())

	if !fileInfo.IsDir() {
		content, err := os.ReadFile(fullPath)
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(string(content), searchStr) {
			fmt.Println(fileInfo.Name(), "contains", searchStr)
		} else {
			fmt.Println(fileInfo.Name(), "doesn't contain", searchStr)
		}
	}
}

func main() {
	searchStr := os.Args[1]
	path := os.Args[2]
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileinfo := range files {
		go grepPath(path, fileinfo, searchStr)
	}

	time.Sleep(time.Second * 3)
}
