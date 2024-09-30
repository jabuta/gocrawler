package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addCurrentUrl(rawCurrentURL string) bool {
	//use bool to determin if you should continue executing the function
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return false
	}
	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return false
	}
	if parsedCurrentURL.Hostname() != cfg.baseURL.Hostname() {
		fmt.Println(parsedCurrentURL.Hostname(), cfg.baseURL.Hostname())
		fmt.Println("4 if hostname")

		if _, ok := cfg.pages[normalizedCurrentURL]; !ok {
			cfg.pages[normalizedCurrentURL] = 1
		} else {
			cfg.pages[normalizedCurrentURL] += 1
		}
		return false
	}
	if _, ok := cfg.pages[normalizedCurrentURL]; ok {
		fmt.Println("5 in url list")
		cfg.pages[normalizedCurrentURL]++
		return false
	}
	cfg.pages[normalizedCurrentURL] = 1
	return true
}
