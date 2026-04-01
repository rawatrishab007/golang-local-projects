package main

import (
	"fmt"
	"net/http"
	"time"
)

type result struct {
	URL      string
	duration time.Duration
	Err      error
}

func checkWebsite(url string, ch chan<- result) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- result{URL: url, Err: err}
		return
	}
	defer resp.Body.Close()
	ch <- result{
		URL:      url,
		duration: time.Since(start),
		Err:      nil,
	}

}
func main() {
	websites := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.reddit.com",
		"student.vedam.org",
		"https://www.crazygames.com/",
	}
	ch := make(chan result, len(websites))
	for _, url := range websites {
		go checkWebsite(url, ch)
	}
	fastest := <-ch
	if fastest.Err != nil {
		fmt.Printf("The fastest website %s has this error %v", fastest.URL, fastest.Err)
	} else {
		fmt.Printf("The fastest website is %s\n", fastest.URL)
		fmt.Printf("Time taken is %v", fastest.duration)
	}
	fmt.Println("\nFull Leaderboard:")
	fmt.Printf("1. %-30s %v\n", fastest.URL, fastest.duration)

	for i := 1; i < len(websites); i++ {
		res := <-ch
		fmt.Printf("%d. %-30s %v\n", i+1, res.URL, res.duration)
	}
}
