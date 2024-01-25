package sitemap_builder

import (
	"fmt"
	"gophercises/link_parser"
	"net/http"
	"net/url"
	"strings"
)

//fetch url
//	extract urls
//	normalize them
//	close body
//filter seen urls
//add urls to the next level
//store urls as slice, maybe transfer them to map later
//use interface to visualize slice of links

func TraverseUrl(rootUrl string, depth int) []link_parser.HTMLLink {
	level := make([]link_parser.HTMLLink, 1)
	level = append(level, link_parser.HTMLLink{Link: rootUrl, Text: "Root"})
	res := make([]link_parser.HTMLLink, 0)
	seen := make(map[string]bool)
	seen[rootUrl] = true
	for len(level) > 0 && depth > 0 {
		newLevel := make([]link_parser.HTMLLink, 0)
		for i := range level {
			fmt.Printf("Current link (%d) :%s%s \n", depth, strings.Repeat("    ", 5-depth), level[i].Link)
			relativeRoot, err := url.Parse(level[i].Link)
			if err != nil {
				fmt.Printf("Could not parse url as !root! %s \n", level[i].Link)
				continue
			}
			fetchedLinks, err := FetchUrl(level[i].Link)
			if err != nil {
				level[i].Text += fmt.Sprintf("Error fetching url %v", err)
				continue
			}
			fetchedLinks = formatLinks(fetchedLinks, relativeRoot, seen)
			newLevel = append(newLevel, fetchedLinks...)
		}
		res = append(res, level...)
		level = newLevel
		depth--
	}
	res = append(res, level...)
	fmt.Println(len(res))
	return res
}

func formatLinks(fetchedLinks []link_parser.HTMLLink, relativeRoot *url.URL, seen map[string]bool) []link_parser.HTMLLink {
	relativeUrls := make([]link_parser.HTMLLink, 0)
	for j := range fetchedLinks {
		parsedUrl, err := url.Parse(fetchedLinks[j].Link)
		if err != nil {
			fetchedLinks[j].Text += fmt.Sprintf("Error parsing url %v", err)
			continue
		}
		if len(parsedUrl.Host) == 0 {
			parsedUrl = relativeRoot.JoinPath(parsedUrl.Path)
		}
		linkStr := strings.TrimRight(parsedUrl.String(), "/")
		if parsedUrl.Host != relativeRoot.Host || parsedUrl.Scheme != relativeRoot.Scheme || seen[linkStr] {
			continue
		}
		fetchedLinks[j].Link = linkStr
		seen[linkStr] = true
		relativeUrls = append(relativeUrls, fetchedLinks[j])
	}
	return relativeUrls
}

var semaphore = make(chan struct{}, 20)

func FetchUrl(url string) ([]link_parser.HTMLLink, error) {
	semaphore <- struct{}{}
	defer func() {<- semaphore}()
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return link_parser.ExtractLinks(res.Body)
}

type UrlSet struct {
	Urls []Url `xml:"url"`
}

type Url struct {
	Location string `xml:"location"`
}
