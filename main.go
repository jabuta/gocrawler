package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

func main() {

	url := getArgs()

	var cfg = config{
		pages:              map[string]int{},
		baseURL:            url,
		concurrencyControl: make(chan struct{}, 100),
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
	}
	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL.String())
	cfg.wg.Wait()
	fmt.Print(cfg.pages)
}

func getArgs() *url.URL {
	if len(os.Args) < 2 {
		log.Fatalf("no website provided")
	} else if len(os.Args) > 2 {
		log.Fatalf("too many arguments provided")
	}
	url, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("Malformed URL: %v", err)
	}
	fmt.Printf("- 'starting crawl'\n- '%s   ", os.Args[1])

	return url
}
