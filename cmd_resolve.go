////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"io"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// resolve

func resolveCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*resolveT)
	// fmt.Printf("[resolve]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.DNSServer, Opts.Port, Opts.Retires, Opts.Verbose =
		rootArgv.DNSServer, rootArgv.Port, rootArgv.Retires, rootArgv.Verbose.Value()
	return cmdResolve(argv.Filei)
}

func cmdResolve(cin io.Reader) error {
	r := NewDnsResolver(false)
	// read cin line by line
	scanner := bufio.NewScanner(cin)
	for scanner.Scan() {
		r.Lookup(scanner.Text())
	}
	return scanner.Err()
}
