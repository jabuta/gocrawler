package main

import (
	"fmt"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()

	if !cfg.addCurrentUrl(rawCurrentURL) {
		return
	}

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
	fmt.Println(urlsToCrawl)
	for _, urlToCrawl := range urlsToCrawl {
		cfg.wg.Add(1)
		go cfg.crawlPage(urlToCrawl)
	}
}
