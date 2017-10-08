////////////////////////////////////////////////////////////////////////////
// Program: dnstools
// Purpose: DNS Tools
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits:
// 				  https://github.com/bogdanovich/dns_resolver/
// 				  https://stackoverflow.com/questions/30043248/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/miekg/dns"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// DnsResolver represents a dns resolver
type DnsResolver struct {
	OptsT
}

// NewDnsResolver initializes DnsResolver at high level.
func NewDnsResolver() *DnsResolver {
	if Opts.DNSServer == "" {
		return NewFromResolvConf("/etc/resolv.conf")
	} else {
		return New()
	}
}

// New initializes DnsResolver.
func New() *DnsResolver {
	return &DnsResolver{Opts}
}

// NewFromResolvConf initializes DnsResolver from resolv.conf like file.
func NewFromResolvConf(path string) *DnsResolver {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		abortOn("Accessing "+path, err)
	}
	config, err := dns.ClientConfigFromFile(path)
	abortOn("Opening "+path, err)

	Opts.DNSServer = config.Servers[0]
	Opts.Port = "53"
	return New()
}

// Lookup prints the provied host and its IP addresses.
func (r *DnsResolver) Lookup(host string) error {
	IPs, err := r.LookupHost(host)
	if err != nil {
		warning(host + ": " + err.Error())
		// return err
	}
	fmt.Printf("%s\t", host)
	for _, ip := range IPs {
		fmt.Printf("%s ", ip)
	}
	fmt.Println()
	return nil
}

// LookupHost returns IP addresses of provied host.
// In case of timeout retries query r.Retires times.
func (r *DnsResolver) LookupHost(host string) ([]net.IP, error) {
	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(dns.Fqdn(host), dns.TypeA)

	Server := net.JoinHostPort(r.DNSServer, r.Port)
	result := []net.IP{}
	in, _, err := c.Exchange(&m, Server)
	for ii := 0; err != nil &&
		strings.HasSuffix(err.Error(), "i/o timeout") &&
		ii < r.Retires; ii++ {
		in, _, err = c.Exchange(&m, Server)
	}
	if err != nil {
		return result, err
	}

	if in != nil && in.Rcode != dns.RcodeSuccess {
		return result, errors.New(dns.RcodeToString[in.Rcode])
	}

	for _, record := range in.Answer {
		if t, ok := record.(*dns.A); ok {
			result = append(result, t.A)
		}
	}
	return result, err
}
