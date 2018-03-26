package crawl

import (
	"fmt"
	"p2b/web_crawler/links"
	"net/http"
	"time"
)

type Crawler interface {
	Crawl(url string, depth int, xtractor links.Xtractor) (map[string]links.Links, error)
}

//generates a string of  sequence of \t chars for print prefixing
func getPrefStr(forDepth int) string {
	prefixstr := make([]rune, forDepth)
	for i := range prefixstr {
		prefixstr[i] = '\t'
	}
	return string(prefixstr)
}



type dfsCrawler struct {
	hcli http.Client
	maxDepth int
}

func NewDfsCrawler(timeoutsec uint, depth int) *dfsCrawler {
	timeout := time.Duration(timeoutsec) * time.Second
	return &dfsCrawler{
		hcli: http.Client{
			Timeout: timeout,
			},
			maxDepth: depth,
	}
}

type results struct{
	url string
	urls links.Links
	err error
	depth int
}

func getLinksFromBody(from string, hcli http.Client, xtr links.Xtractor) (links links.Links, err error) {
	var r *http.Response
	r, err = hcli.Get(from)
	if err != nil{
		return nil, err
	}
	defer r.Body.Close()
	// perform the http.get() on the url. check for error
	// Using the Xtractor, extract all the links and return the links and error
	myLinks, err := xtr.Xtract(r.Body)
	// Remember the close the body of the response before closing

	return myLinks, err

}
func (c *dfsCrawler) Crawl(url string, depth int, xtr links.Xtractor) (map[string]bool, error) {
	var err error
	//create a map to store all the links extracted from one url.
	fetched := make(map[string]bool)

	if depth > 1 {
		//_, ok := fetched[url]
		if !fetched[url]{
			urls, err := getLinksFromBody(url, c.hcli, xtr)
			if err != nil{
				fmt.Println("error: failed to crawl "+url+"")
			}
			for _, u := range urls{
				fmt.Printf("%s %s\n", getPrefStr((c.maxDepth+1) - depth), u.String())
				c.Crawl(u.String(), depth - 1, xtr)
				fetched[u.String()] = true
			}
		}
	}
	return fetched, err
	//get links from body using the function getLinksFromBody() defined above.
	// Using recursion or go routines, implement the depth first search.

}
