package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {

	cfg := initCfg()

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL.String())
	cfg.wg.Wait()

	for normalizedURL, qty := range cfg.pages {
		fmt.Printf("%s - %d\n", normalizedURL, qty)
	}
	cfg.printReport()
}

func initCfg() config {
	if len(os.Args) < 4 {
		log.Fatalf("no website provided")
	} else if len(os.Args) > 4 {
		log.Fatalf("too many arguments provided")
	}
	url, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("Malformed URL: %v", err)
	}
	fmt.Printf("- 'starting crawl'\n- '%s   ", os.Args[1])

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Second argument is not an int")
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal("Third argument is not an int")
	}
	if maxPages < 1 || maxConcurrency < 1 {
		log.Fatal("Second and third arguments need to be a positive integer")
	}

	var cfg = config{
		pages:              map[string]int{},
		baseURL:            url,
		concurrencyControl: make(chan struct{}, maxConcurrency),
		mu:                 &sync.RWMutex{},
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	return cfg
}
