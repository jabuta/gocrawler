package main

import (
	"fmt"
	"sort"
)

type link struct {
	url string
	qty int
}

type byLinkQty []link

func (m byLinkQty) Len() int           { return len(m) }
func (m byLinkQty) Less(i, j int) bool { return m[i].qty > m[j].qty }
func (m byLinkQty) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (cfg config) printReport() {

	i := 0
	pageInt := make([]link, len(cfg.pages))
	for url, qty := range cfg.pages {
		pageInt[i] = link{url: url, qty: qty}
		i++
	}
	sort.Sort(byLinkQty(pageInt))

	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, cfg.baseURL)

	for _, url := range pageInt {
		fmt.Printf("Found %d internal links to %s\n", url.qty, url.url)
	}

}
