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
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	fmt.Println(padding, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
