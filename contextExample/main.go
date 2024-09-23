package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Response struct to store data
type Response struct {
	value int
	err   error
}

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 10

	val, err := fetchAPIData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value:", val)
	fmt.Println("Time Taken:", time.Since(start))
}

func fetchAPIData(ctx context.Context, userID int) (int, error) {
	// Set a timeout of 100ms
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	// Channel to receive response
	respch := make(chan Response)

	// Fetch data in a goroutine
	go func() {
		val, err := ThirdPartyAPI()
		respch <- Response{
			value: val,
			err:   err,
		}
		close(respch) // Close the channel when done
	}()

	// Wait for either the response or timeout
	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("fetching from 3rd party API took too long")
	case resp := <-respch:
		return resp.value, resp.err
	}
}

// Simulated third-party API with a 500ms delay
func ThirdPartyAPI() (int, error) {
	time.Sleep(time.Millisecond * 500) // Simulating delay
	return 200, nil
}
