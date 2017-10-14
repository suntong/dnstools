////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

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
	Opts.NxDomainErr = true
	if argv.Dig {
		cmdDig(ctx.Args()[0])
	}
	Opts.Stop = argv.Stop
	return cmdProbe(ctx.Args(), argv.Raw)
}

func cmdProbe(hosts []string, raw bool) error {
	r := NewDnsResolver(false)
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

// cmdDig, dig mode, output only IP without domain name
// exit code 1 if error
func cmdDig(host string) {
	r := NewDnsResolver(true)
	e := r.Lookup(host)
	if e != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
