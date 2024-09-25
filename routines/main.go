package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Response struct {
	Link   string
	Status int
}

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.in",
		"http://amuserrr.in",
	}
	var wg sync.WaitGroup

	c := make(chan Response)

	wg.Add(len(links))

	for _, link := range links {
		go func(link string) {
			defer wg.Done()
			status := checkLink(link)
			c <- Response{Link: link, Status: status}
		}(link)
	}

	go func() {
		wg.Wait()
		close(c) // Close the channel when done
	}()

	// Collect results from the channel
	for response := range c {
		if response.Status != http.StatusOK {
			fmt.Printf("Error checking link %s: %v\n", response.Link, response.Status)
		} else {
			fmt.Printf("Link %s is UP (Status: %d)\n", response.Link, response.Status)
		}
	}
}

func checkLink(link string) int {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error checking link %s: %v\n", link, err)
		return -1
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
