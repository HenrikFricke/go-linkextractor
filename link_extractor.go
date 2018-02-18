package linkextractor

import (
	"errors"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

//WebsiteLinks is a list of links on a website page
type WebsiteLinks struct {
	URL   *url.URL
	Links []*url.URL
}

func sanitizeLink(link string, websiteURL *url.URL) (*url.URL, error) {
	linkURL, err := url.Parse(link)

	if err != nil {
		return nil, err
	}

	if linkURL.IsAbs() && linkURL.Scheme != "http" && linkURL.Scheme != "https" {
		return nil, errors.New("Unsupported scheme")
	}

	if !linkURL.IsAbs() {
		resolvedURL := websiteURL.ResolveReference(linkURL)
		return resolvedURL, nil
	}

	return linkURL, nil
}

//ExtractLinks scrapes a specific URL and returns all links on the page
func ExtractLinks(URL *url.URL) (*WebsiteLinks, error) {
	httpRes, err := http.Get(URL.String())

	if err != nil {
		return nil, err
	}

	z := html.NewTokenizer(httpRes.Body)
	resolvedURL := httpRes.Request.URL
	links := []*url.URL{}

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return &WebsiteLinks{resolvedURL, links}, nil
		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						link, err := sanitizeLink(a.Val, resolvedURL)

						if err != nil {
							break
						}

						links = append(links, link)
					}
				}
			}
		}
	}
}
