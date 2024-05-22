package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// crawlWeb function accepts an input channel for sending URLs,
// a context for cancellation, and returns another channel for crawling results.
func crawlWeb(ctx context.Context, urls <-chan string, done <-chan struct{}) <-chan *string {
	results := make(chan *string)

	go func() {
		defer close(results)

		var wg sync.WaitGroup
		defer wg.Wait()

		worker := make(chan struct{}, 8) // Limiting to 8 goroutines querying pages in parallel

		for url := range urls {
			select {

			case <-ctx.Done():
				return // Stop if cancellation signal received

			case worker <- struct{}{}:
				wg.Add(1)

				go func(url string) {
					defer wg.Done()
					defer func() { <-worker }()
					body, err := fetchURL(url)
					if err != nil {
						fmt.Printf("Error fetching %s: %s\n", url, err)
						return
					}
					select {
					case <-ctx.Done():
						return // Stop if cancellation signal received
					case results <- body:
					}
				}(url)

			}
		}

	}()

	return results
}

// fetchURL function fetches the body of a URL and returns it as a string pointer.
func fetchURL(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyStr := string(body)
	return &bodyStr, nil
}

func main() {
	urls := make(chan string)
	done := make(chan struct{})

	// Create a context with cancellation ability
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Graceful shutdown on interrupt signal (Ctrl+C)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		fmt.Println("\nInterrupt signal received. Stopping crawling...")
		cancel()
		close(done)
	}()

	// Start crawling
	go func() {
		for i := 0; i < 10; i++ {
			urls <- "https://example.com"
		}
		// Add more URLs to crawl if needed
		close(urls)
	}()

	// Receive crawling results
	results := crawlWeb(ctx, urls, done)

	// Process crawling results
	for body := range results {
		fmt.Println(*body)
	}
}
