package links

import (
	"fmt"
	"io"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

type link struct {
	name string
}

type Links []link

func NewLinks(from ...string) Links {
	ret := make(Links, len(from))
	for i := range from {
		nl, err := NewLink(from[i])
		if err != nil {
			fmt.Printf("error making links:%s", err)
			break
		}
		ret[i] = nl
	}
	return ret
}

func (l link) String() string {
	return fmt.Sprintf("%s", l.name)
}

func NewLink(a string) (link, error) {
	//validate the URL before accepting it
	_, err := url.Parse(a)
	if err != nil {
		return link{}, err
	}
	return link{
		name: a,
	}, nil
}

type Xtractor interface {
	Xtract(io.Reader) ([]link, error)
}

type Closer interface {
	Close()
}

type substringXtractor struct {
}

func NewSubstringXtractor() *substringXtractor {
	return &substringXtractor{}
}

type parserXtractor struct {
}

func NewParserXtractor() *parserXtractor {
	return &parserXtractor{}
}

// only gets http and https links
func (p *parserXtractor) Xtract(in io.Reader) (ret []link, err error) {
	var (
		n     *html.Node
		inner func(*html.Node)
	)
	if n, err = html.Parse(in); err != nil {
		return
	}
	inner = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					//only attempt to parse http and https (external) urls
					if strings.HasPrefix(a.Val, "http") || strings.HasPrefix(a.Val, "https") {
						nl, err := NewLink(a.Val)
						if err != nil {
							return
						}
						ret = append(ret, nl)
					}
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			inner(c)
		}
	}
	inner(n)
	return
}

func (r *substringXtractor) Xtract(in io.Reader) ([]link, error) {
	nl, err := NewLink("http://dont.do.that.type.of/parsing")
	return []link{nl}, err
}

