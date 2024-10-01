package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.RWMutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()

	if func() bool {
		cfg.mu.RLock()
		defer cfg.mu.RUnlock()
		return !(len(cfg.pages) < cfg.maxPages)
	}() {
		return
	}

	if !cfg.addCurrentUrl(rawCurrentURL) {
		return
	}
	fmt.Printf("crawling %s\n", rawCurrentURL)
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Couldn't crawl %s, error: %v\n", rawCurrentURL, err)
		return
	}
	urlsToCrawl, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Couldn't parse html in %s,\n error: %v\n", rawCurrentURL, err)
		return
	}
	for _, urlToCrawl := range urlsToCrawl {
		cfg.wg.Add(1)
		go cfg.crawlPage(urlToCrawl)
	}
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

		if _, ok := cfg.pages[normalizedCurrentURL]; !ok {
			cfg.pages[normalizedCurrentURL] = 1
		} else {
			cfg.pages[normalizedCurrentURL] += 1
		}
		return false
	}
	if _, ok := cfg.pages[normalizedCurrentURL]; ok {
		cfg.pages[normalizedCurrentURL]++
		return false
	}
	cfg.pages[normalizedCurrentURL] = 1
	return true
}
