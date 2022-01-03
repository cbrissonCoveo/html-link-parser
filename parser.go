package link

import "io"

// Link is a link (<a href="...">) in an HTML page
type Link struct {
	Href string
	Text string
}

// Parse takes an HTML page and returns
// a slice of links or an error.
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
