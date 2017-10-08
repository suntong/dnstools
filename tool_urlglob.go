////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"regexp"
	"strings"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const (
	URLPlain = iota
	URLSet
	URLRange
)

// The URLPattern holds the elements from url glob patterns.
type URLPattern struct {
	ptype   int // pattern type
	pattern []string
}

type URLGlob struct {
	URL     string
	urlGlob []URLPattern
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	reURLGlob = regexp.MustCompile(`[[{].*?[]}]`)
	reRange   = regexp.MustCompile(`\[(.*?)-(.*?)\]`)
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

func NewURLGlob(hosts string) *URLGlob {
	return &URLGlob{URL: hosts}
}

// Parse will parse URL into urlGlob
func (g *URLGlob) Parse() {
	// https://play.golang.org/p/KEiq7__4Ce
	indices := reURLGlob.FindAllStringSubmatchIndex(g.URL, -1)

	// ii is char pointer into g.URL; jj is index of slice indices
	ii, jj := 0, 0
	for ii < len(g.URL) && jj < len(indices) {
		var up URLPattern
		til := indices[jj][0]
		if ii < til {
			// plain text exist before next match
			up.ptype = URLPlain
			up.pattern = []string{g.URL[ii:til]}
			ii = til
		} else {
			// here ii must == til, i.e., ii pointing to a match
			end := indices[jj][1]
			switch g.URL[til] { // use first char to determine type
			case '{':
				// set expression, from the opening '{'	til the next closing '}'
				// with ','-separated elements in between
				// e.g., site.{one,two,three}.com
				up.ptype = URLSet
				up.pattern = strings.Split(g.URL[til+1:end-1], ",")
			case '[':
				/* range expression, with the opening '[', and
				   - char range: e.g. "a-z]", "B-Q]"
				   - num range: e.g. "0-9]", "17-2000]"
				   - num range with leading zeros: e.g. "001-999]"
				   until the next ']'
				*/
				up.ptype = URLRange
				up.pattern = reRange.FindStringSubmatch(g.URL[til:end])
			}
			ii = end
			jj++
		}
		g.urlGlob = append(g.urlGlob, up)
	}
}
