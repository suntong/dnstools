////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// probe

func probeCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*probeT)
	// fmt.Printf("[probe]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.DNSServer, Opts.Port, Opts.Retires, Opts.Verbose =
		rootArgv.DNSServer, rootArgv.Port, rootArgv.Retires, rootArgv.Verbose.Value()
	Opts.Stop = argv.Stop

	// g := NewURLGlob("site.{one,two,three}[1-100]-{one,two,three}[2-20].com")
	g := NewURLGlob(ctx.Args()[0])
	g.Parse()
	fmt.Println(g.urlGlob)
	fmt.Println(g.GetURLs(0))
	return nil

	Opts.NxDomainErr = true
	r := NewDnsResolver()
	r.Lookup(ctx.Args()[0])

	return nil
}
