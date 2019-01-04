////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cloverstd/tcping/ping"
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

	timeoutDuration, err := time.ParseDuration(argv.Timeout)
	clis.AbortOn("Parsing Timeout", err)
	intervalDuration, err := time.ParseDuration(argv.Interval)
	clis.AbortOn("Parsing Interval", err)

	args := ctx.Args()
	host := args[0]
	port := 80
	if len(args) == 2 {
		port, err = strconv.Atoi(args[1])
		clis.AbortOn("Port should be integer", err)
	}

	schema := ping.TCP.String()
	clis.Verbose(2, "Schema: %s\n", schema)
	protocol, _ := ping.NewProtocol(schema)

	target := ping.Target{
		Timeout:  timeoutDuration,
		Interval: intervalDuration,
		Host:     host,
		Port:     port,
		Counter:  argv.Counter,
		Protocol: protocol,
	}

	//return nil
	return DoTcping(target)
}

//
// DoTcping implements the business logic of command `tcping`
func DoTcping(target ping.Target) error {
	fmt.Fprintf(os.Stderr, "%s v%s. tcping - Ping over tcp\n", progname, version)
	fmt.Fprintln(os.Stderr, "Copyright (C) 2019, Tong Sun\n")

	pinger := ping.NewTCPing()
	pinger.SetTarget(&target)
	pingerDone := pinger.Start()

	var sigs chan os.Signal
	select {
	case <-pingerDone:
		break
	case <-sigs:
		break
	}

	r := pinger.Result()
	clis.Verbose(1, "%s", r)

	if r.SuccessCounter == 0 {
		return errors.New("Ping failed: Host down")
	}

	return nil
}
