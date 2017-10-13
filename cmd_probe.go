////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/mkideal/cli"
	"github.com/suntong/curlurl"
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
	Opts.NxDomainErr = true
	return cmdProbe(ctx.Args(), argv.Raw)
}

func cmdProbe(hosts []string, raw bool) error {
	r := NewDnsResolver()
	for _, hp := range hosts {
		g := curlurl.NewURLGlob(hp).Parse(abortOn) // parse the host parameters
		for _, h := range g.GetURLs(0) {
			if raw {
				fmt.Println(h)
			} else {
				r.Lookup(h)
			}
		}
	}
	return nil
}
