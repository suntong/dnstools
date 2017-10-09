////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func ExampleURLGlob_getURLs_output() {
	getURLsTest("google.ca")
	getURLsTest("site.{one,two,three}.com")
	getURLsTest("site.{one,two,three}")
	getURLsTest("site[1-3].com")
	getURLsTest("site.{one,two,three}[1-3].com")
	getURLsTest("site.{one,two,three}[1-3]")
	getURLsTest("{one,two,three}[1-3]")
	getURLsTest("site.{one,two,three}[1-3]-{one,two,three}[8-10].com")
	// Output:
	// -
	// google.ca
	// [{0 [google.ca]}]
	// [google.ca]
	// -
	// site.{one,two,three}.com
	// [{0 [site.]} {1 [one two three]} {0 [.com]}]
	// [site.one.com site.two.com site.three.com]
	// -
	// site.{one,two,three}
	// [{0 [site.]} {1 [one two three]}]
	// [site.one site.two site.three]
	// -
	// site[1-3].com
	// [{0 [site]} {2 [1 2 3]} {0 [.com]}]
	// [site1.com site2.com site3.com]
	// -
	// site.{one,two,three}[1-3].com
	// [{0 [site.]} {1 [one two three]} {2 [1 2 3]} {0 [.com]}]
	// [site.one1.com site.one2.com site.one3.com site.two1.com site.two2.com site.two3.com site.three1.com site.three2.com site.three3.com]
	// -
	// site.{one,two,three}[1-3]
	// [{0 [site.]} {1 [one two three]} {2 [1 2 3]}]
	// [site.one1 site.one2 site.one3 site.two1 site.two2 site.two3 site.three1 site.three2 site.three3]
	// -
	// {one,two,three}[1-3]
	// [{1 [one two three]} {2 [1 2 3]}]
	// [one1 one2 one3 two1 two2 two3 three1 three2 three3]
	// -
	// site.{one,two,three}[1-3]-{one,two,three}[8-10].com
	// [{0 [site.]} {1 [one two three]} {2 [1 2 3]} {0 [-]} {1 [one two three]} {2 [8 9 10]} {0 [.com]}]
	// [site.one1-one8.com site.one1-one9.com site.one1-one10.com site.one1-two8.com site.one1-two9.com site.one1-two10.com site.one1-three8.com site.one1-three9.com site.one1-three10.com site.one2-one8.com site.one2-one9.com site.one2-one10.com site.one2-two8.com site.one2-two9.com site.one2-two10.com site.one2-three8.com site.one2-three9.com site.one2-three10.com site.one3-one8.com site.one3-one9.com site.one3-one10.com site.one3-two8.com site.one3-two9.com site.one3-two10.com site.one3-three8.com site.one3-three9.com site.one3-three10.com site.two1-one8.com site.two1-one9.com site.two1-one10.com site.two1-two8.com site.two1-two9.com site.two1-two10.com site.two1-three8.com site.two1-three9.com site.two1-three10.com site.two2-one8.com site.two2-one9.com site.two2-one10.com site.two2-two8.com site.two2-two9.com site.two2-two10.com site.two2-three8.com site.two2-three9.com site.two2-three10.com site.two3-one8.com site.two3-one9.com site.two3-one10.com site.two3-two8.com site.two3-two9.com site.two3-two10.com site.two3-three8.com site.two3-three9.com site.two3-three10.com site.three1-one8.com site.three1-one9.com site.three1-one10.com site.three1-two8.com site.three1-two9.com site.three1-two10.com site.three1-three8.com site.three1-three9.com site.three1-three10.com site.three2-one8.com site.three2-one9.com site.three2-one10.com site.three2-two8.com site.three2-two9.com site.three2-two10.com site.three2-three8.com site.three2-three9.com site.three2-three10.com site.three3-one8.com site.three3-one9.com site.three3-one10.com site.three3-two8.com site.three3-two9.com site.three3-two10.com site.three3-three8.com site.three3-three9.com site.three3-three10.com]

}

func getURLsTest(hostn string) {
	fmt.Printf("-\n%s\n", hostn)
	g := NewURLGlob(hostn)
	g.Parse()
	fmt.Println(g.urlGlob)
	fmt.Println(g.GetURLs(0))
}
