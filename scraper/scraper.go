// scraper.go

package scraper

import (
	"fmt"
	"net"
	"net/url"

	"github.com/gocolly/colly"
)

// Article ..
type Article struct {
	headline    string
	description string
	link        string
	author      string
	image       string
	date        int32
}

// InitScraper ...
func InitScraper(s string) error {
	url, _ := parseURL(s)
	fmt.Println("Crawling on: ", url)
	scraper := colly.NewCollector(
		colly.AllowedDomains("bt.dk", "www.bt.dk"),
	)

	// On every a element which has href attribute call callback
	scraper.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		scraper.Visit(e.Request.AbsoluteURL(link))
	})

	scraper.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	return scraper.Visit(url)
}

func (a Article) copyContent() {}

func parseURL(s string) (string, string) {
	u, err := url.Parse(s)
	host, _, _ := net.SplitHostPort(u.Host)
	fmt.Println("Parsed host: ", host)
	if err != nil {
		panic(err)
	}
	return s, host
}
