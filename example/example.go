package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/HenrikFricke/go-linkextractor"
)

func main() {
	inputURL := os.Getenv("URL")
	parsedURL, _ := url.Parse(inputURL)
	res, _ := linkextractor.ExtractLinks(parsedURL)

	fmt.Println("URL: ", res.URL)
	fmt.Println("")
	fmt.Println("Links:")

	for _, link := range res.Links {
		fmt.Println(link)
	}

	fmt.Println("")
	fmt.Println("Number of links: ", len(res.Links))
}
