package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"p2b/web_crawler/links"
	"p2b/web_crawler/crawl"
)

var (
	flagUrl         string
	linkSearchDepth int
	flagTimeoutSecs uint
	sitesCrawled    int
	t1              time.Time
	hcli            http.Client
)

func init() {
	// Import as explained here
	flagUrl = ""
	linkSearchDepth = 0
	flagTimeoutSecs = 100
	sitesCrawled = 0
	t1 = time.Now()
}

func main() {
	// Parse the flag here
	flagUrl := flag.String("url", "","url")
	linkSearchDepth := flag.Int("depth", 0 , "depth")
	flagTimeoutSecs := flag.Uint("timeout", 5, "timeout")
	flag.Parse()
	if *flagUrl == ""{
		panic("Url not provided")
	}
	if *linkSearchDepth == 0{
		panic("Depth not provided")
	}

	// Create a new ParserXtractor
	xtr := links.NewParserXtractor()
	// Create a new DFS Crawler
	dfs := crawl.NewDfsCrawler(*flagTimeoutSecs)
	// Create a Time variable using the time package and record the time
	t1 = time.Now()
	// Run the Crawl function and print the length of the Crawled output and the time taken
	linkMap, err := dfs.Crawl(*flagUrl, *linkSearchDepth, xtr)
	sitesCrawled = len(linkMap)
	d := time.Since(t1)
	fmt.Printf("Crawled: %d time taken: %f err: %s\n", sitesCrawled, d.Seconds(), err)
}
