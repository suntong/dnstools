////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// tcping

func tcpingCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*tcpingT)
	clis.Setup(progname, rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	Opts.DNSServer, Opts.Port, Opts.Retires, Opts.Verbose =
		rootArgv.DNSServer, rootArgv.Port, rootArgv.Retires, rootArgv.Verbose.Value()
	//return nil
	return DoTcping()
}

//
// DoTcping implements the business logic of command `tcping`
func DoTcping() error {
	fmt.Fprintf(os.Stderr, "%s v%s. tcping - Ping over tcp\n", progname, version)
	fmt.Fprintln(os.Stderr, "Copyright (C) 2019, Tong Sun\n")
	return nil
}
