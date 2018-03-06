package crawl

import (
	"fmt"
	"p2b/web_crawler/links"
	"net/http"
	"time"
	"testing"
)

type Crawler interface {
	Crawl(url string, depth int, xtractor links.Xtractor) (map[string]links.Links, error)
}

type result struct{
	url string
	urls links.Links
	err error
	depth int
}

//generates a string of  sequence of \t chars for print prefixing
func getPrefStr(forDepth int) string {
	prefixstr := make([]rune, forDepth)
	for i := range prefixstr {
		prefixstr[i] = '\t'
	}
	return string(prefixstr)
}

func getLinksFromBody(from string, hcli http.Client, xtr links.Xtractor) (links links.Links, err error) {
	var r *http.Response
	r, err = hcli.Get(from)
	if err != nil{
		panic(err)
	}
	defer r.Body.Close()
	if err != nil{
		panic(err)
	}
	// perform the http.get() on the url. check for error
	// Using the Xtractor, extract all the links and return the links and error
	myLinks, err := xtr.Xtract(r.Body)
	// Remember the close the body of the response before closing

	return myLinks, err

}


type dfsCrawler struct {
	hcli http.Client
}

func NewDfsCrawler(timeoutsec uint) *dfsCrawler {
	tdur := time.Duration(time.Duration(timeoutsec) * time.Second)
	return &dfsCrawler{
		hcli: http.Client{
			Timeout: tdur,
			},

	}
}

func (c *dfsCrawler) Crawl(url string, depth int, xtr links.Xtractor) (map[string]links.Links, error) {
	var (
		err     error
	)
	//create a map to store all the links extracted from one url.
	fetched := make(map[string]bool)
	results := make(chan *result)
	recurse := func(url string, depth int){
		urls, err := getLinksFromBody(url, c.hcli, xtr)
		results <- &result{url, urls, err, depth}
	}


	go recurse(url, depth)
	fetched[url] = true

	for fetching := 1; fetching > 0; fetching--{
		res := <- results

		if res.err != nil{
			fmt.Println(res.err)
			continue
		}


	}
	//get links from body using the function getLinksFromBody() defined above.
	// Using recursion or go routines, implement the depth first search.


	return retlinks, err
}
