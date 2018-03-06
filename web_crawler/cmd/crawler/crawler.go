package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"os"
	"strconv"
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
	hcli = *(http.DefaultClient)
}

func main() {
	args := os.Args[1:]
	// Parse the flag here
	flagUrl = *(flag.String("url", args[1], "url"))
	linkSearchDepth, _ = strconv.Atoi(args[2])
	linkSearchDepth = *(flag.Int("depth", linkSearchDepth , "depth"))
	tmp, _ := strconv.ParseUint(args[3], 10, 32)
	flagTimeoutSecs = uint(tmp)
	flagTimeoutSecs = *(flag.Uint("timeout", flagTimeoutSecs, "timeout"))
	// Create a new ParserXtractor
	xtr := links.NewParserXtractor()
	// Create a new DFS Crawler
	dfs := crawl.NewDfsCrawler(flagTimeoutSecs)
	// Create a Time variable using the time package and record the time
	t1 = time.Now()
	// Run the Crawl function and print the length of the Crawled output and the time taken
	linkMap, _ := dfs.Crawl(flagUrl, linkSearchDepth, xtr)
	sitesCrawled = len(linkMap)
	d := time.Since(t1)
	fmt.Println("Time taken: %d\n Length of output: %d", d.String(), sitesCrawled)
}
