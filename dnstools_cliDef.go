////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	//  	"fmt"
	//  	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// dnstools

type rootT struct {
	cli.Helper
	DNSServer string      `cli:"H,host" usage:"dns server IP addr"`
	Port      string      `cli:"p,port" usage:"dns server port" dft:"53"`
	Retires   int         `cli:"retry" usage:"retry this many times when dns query times out" dft:"2"`
	Verbose   cli.Counter `cli:"v,verbose" usage:"verbose mode (multiple -v options increase the verbosity.)"`
}

var root = &cli.Command{
	Name:   "dnstools",
	Desc:   "DNS Tools\nVersion " + version + " built on " + date,
	Text:   "Tool for DNS inquiry",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     dnstools,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	DNSServer	string
//  	Port	string
//  	Retires	int
//  	Verbose	cli.Counter
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "dnstools"
//          version   = "0.1.0"
//          date = "2019-01-03"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,
//  		cli.Tree(resolveDef),
//  		cli.Tree(probeDef),
//  		cli.Tree(tcpingDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

//  func dnstools(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// resolve

//  func resolveCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*resolveT)
//  	clis.Setup(progname, rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.DNSServer, Opts.Port, Opts.Retires, Opts.Verbose, Opts.Verbose =
//  		rootArgv.DNSServer, rootArgv.Port, rootArgv.Retires, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	//return nil
//  	return DoResolve()
//  }
//
// DoResolve implements the business logic of command `resolve`
//  func DoResolve() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. resolve - Resolve from given domain list to IP\n", progname, version)
//  	fmt.Fprintln(os.Stderr, "Copyright (C) 2019, Tong Sun\n")
//  	return nil
//  }

type resolveT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"domain list file (mandatory, read stdin if no file given)"`
}

var resolveDef = &cli.Command{
	Name: "resolve",
	Desc: "Resolve from given domain list to IP",
	Text: "Usage:\n  dnstools resolve -i [/TMP/LISTF]",
	Argv: func() interface{} { return new(resolveT) },
	Fn:   resolveCLI,

	NumOption: cli.AtLeast(1),
}

////////////////////////////////////////////////////////////////////////////
// probe

//  func probeCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*probeT)
//  	clis.Setup(progname, rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.DNSServer, Opts.Port, Opts.Retires, Opts.Verbose, Opts.Verbose =
//  		rootArgv.DNSServer, rootArgv.Port, rootArgv.Retires, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	//return nil
//  	return DoProbe()
//  }
//
// DoProbe implements the business logic of command `probe`
//  func DoProbe() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. probe - Probe the given domain set to IP\n", progname, version)
//  	fmt.Fprintln(os.Stderr, "Copyright (C) 2019, Tong Sun\n")
//  	return nil
//  }

type probeT struct {
	Raw  bool `cli:"raw" usage:"output raw domain set globing result without IP"`
	Dig  bool `cli:"dig" usage:"dig mode, output only IP without domain name"`
	Stop int  `cli:"stop" usage:"stop probing after this many errors" dft:"3"`
}

var probeDef = &cli.Command{
	Name: "probe",
	Desc: "Probe the given domain set to IP",
	Text: "Usage:\n  dnstools probe DOMAIN_SET",
	Argv: func() interface{} { return new(probeT) },
	Fn:   probeCLI,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}

////////////////////////////////////////////////////////////////////////////
// tcping

//  func tcpingCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*tcpingT)
//  	clis.Setup(progname, rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.DNSServer, Opts.Port, Opts.Retires, Opts.Verbose, Opts.Verbose =
//  		rootArgv.DNSServer, rootArgv.Port, rootArgv.Retires, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	//return nil
//  	return DoTcping()
//  }
//
// DoTcping implements the business logic of command `tcping`
//  func DoTcping() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. tcping - Ping over tcp\n", progname, version)
//  	fmt.Fprintln(os.Stderr, "Copyright (C) 2019, Tong Sun\n")
//  	return nil
//  }

type tcpingT struct {
	Counter   int    `cli:"c,counter" usage:"ping counter" dft:"3"`
	Timeout   string `cli:"T,timeout" usage:"connect timeout, units are 'ms', 's'" dft:"3s"`
	Interval  string `cli:"I,interval" usage:"ping interval, units are 'ms', 's'" dft:"2s"`
	DnsServer string `cli:"dns" usage:"Use the specified dns resolve server"`
}

var tcpingDef = &cli.Command{
	Name: "tcping",
	Desc: "Ping over tcp",
	Text: "Usage:\n  dnstools tcping domain_name/IP [port]",
	Argv: func() interface{} { return new(tcpingT) },
	Fn:   tcpingCLI,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}
