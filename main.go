package main

import (
	"fmt"
	"log"

	"github.com/byblix/webscraper/scraper"
)

func main() {
	fmt.Println("Running scraper")
	err := scraper.InitScraper("https://www.bt.dk/")
	if err != nil {
		log.Fatalln(err)
	}
}
