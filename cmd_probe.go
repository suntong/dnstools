////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
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
	Opts.NxDomainErr = true
	return cmdProbe(ctx.Args())
}

func cmdProbe(hosts []string) error {
	r := NewDnsResolver()
	for _, hp := range hosts {
		g := NewURLGlob(hp).Parse() // parse the host parameters
		for _, h := range g.GetURLs(0) {
			// print(h, "\n")
			r.Lookup(h)
		}
	}
	return nil
}
