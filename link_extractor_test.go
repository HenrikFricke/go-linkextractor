package linkextractor

import (
	"net/url"
	"reflect"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestLinkExtractor(t *testing.T) {
	defer gock.Off()

	testURL := "http://web.site/"
	parsedURL, _ := url.Parse(testURL)

	gock.New(testURL).
		Get("/").
		Reply(200).
		BodyString(`
			<!DOCTYPE html>
			<html>
				<body>
					<header>
						<h1>My Website</h1>
					</header>
					<main>
						<a href="/subpage">Subpage</a>
						<a href="http://web.site/subpage2">Subpage 2</a>
						<a href="https://google.com">Visit Google</a>
						<a href="mailto:test@website.com">Mail me</a>
						<a href="javascript: alert('nice');">Mail me</a>
					</main>
				</body>
			</html>
			`)

	link1, _ := url.Parse("http://web.site/subpage")
	link2, _ := url.Parse("http://web.site/subpage2")
	link3, _ := url.Parse("https://google.com")

	expectedWebsiteLinks := WebsiteLinks{
		URL: parsedURL,
		Links: []*url.URL{
			link1,
			link2,
			link3}}

	res, err := ExtractLinks(testURL)

	if err != nil {
		t.Error(err)
	}

	if res.URL.String() != expectedWebsiteLinks.URL.String() {
		t.Error("URL in response not equal.")
	}

	if !reflect.DeepEqual(res.Links, expectedWebsiteLinks.Links) {
		t.Error("List of links not equal.", res.Links, expectedWebsiteLinks.Links)
	}
}
