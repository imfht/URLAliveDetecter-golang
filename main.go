package URLAliveDetecter

import (
	"github.com/parnurzeal/gorequest"
	"fmt"
	"sync"
	"time"
	"os"
	"log"
	"bufio"
)

func doRequest(url string) {
	_, _, errs := gorequest.New().Get(url).End()
	if errs != nil {
		fmt.Print(url)
	}
}

func Worker(urls []string) {
		var wg sync.WaitGroup
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			_, _, errs := gorequest.New().Timeout(2 * time.Second).Head(url).End()
			if errs == nil {
				fmt.Println(url)
			}
		}(url)
	}
		wg.Wait()
}

func WorkFile(path string) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var urls = []string{}
	for sc.Scan() {
		url := sc.Text()
		urls = append(urls, url)
		if len(urls) == 200 {
			Worker(urls)
      fmt.Println("work on 200 urls chunk")
			urls = []string{}
		}
	}
  fmt.Println("work on ",len(urls),"urls")
	Worker(urls)
}
