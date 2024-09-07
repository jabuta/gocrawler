package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	nodes, err := html.Parse(htmlReader)
	if err != nil {
		return []string{}, err
	}
	links := []string{}
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "button") {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href := attr.Val
					if !strings.Contains(href, rawBaseURL) {
						href = rawBaseURL + href
					}
					links = append(links, href)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = n.NextSibling {
			f(c)
		}
	}
	f(nodes)
	return []string{}, nil
}
