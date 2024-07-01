package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

func wordCounter(url string, frequency safeMap) {
	// Download the web-page from given url
	resp, _ := http.Get(url)
	if resp.StatusCode != 200 {
		panic(url + "returned status code" + resp.Status)
	}
	// Closes the response at the end of the function
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	for _, word := range wordRegex.FindAllString(string(body), -1) {
		wordLower := strings.ToLower(word)
		frequency.mu.Lock()
		frequency.registry[wordLower] += 1
		frequency.mu.Unlock()
	}

	fmt.Println("completed:", url)
}

// Defining type of safe-map for avoiding race condition
type safeMap struct {
	registry map[string]int
	mu       *sync.Mutex
}

func main() {
	// Initialixing our safemap
	frequency := safeMap{
		registry: make(map[string]int),
		mu:       &sync.Mutex{},
	}

	for i := 1000; i <= 1030; i++ {
		// Iterate from document id 1000 to 1030 to download 31 docs
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)

		// Debugging
		fmt.Printf("downloading %s \n", url)

		go wordCounter(url, frequency)
	}

	// Waiting for goroutines to finished
	time.Sleep(10 * time.Second)

	// Printing result
	for word, word_frequency := range frequency.registry {
		fmt.Printf("%s -> %d \n", word, word_frequency)
	}
}
