package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link is a link (<a href="...">) in an HTML page
type Link struct {
	Href string
	Text string
}

// Parse takes an HTML page and returns
// a slice of links or an error.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	return nil, nil
}

// linkNodes return a list of <a> nodes
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret

}
