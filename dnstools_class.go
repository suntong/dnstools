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
	Dig bool // dig mode, output only IP
}

// NewDnsResolver initializes DnsResolver at high level.
func NewDnsResolver(dig bool) (r *DnsResolver) {
	if Opts.DNSServer == "" {
		r = NewFromResolvConf("/etc/resolv.conf")
	} else {
		r = New()
	}
	r.Dig = dig
	return
}

// New initializes DnsResolver.
func New() *DnsResolver {
	return &DnsResolver{OptsT: Opts}
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
		// NxDomainErr==true when NXDOMAIN considered error
		if r.NxDomainErr && err.Error() == "NXDOMAIN" {
			return err
		}
		warning(host + ": " + err.Error())
	}
	if !r.Dig {
		fmt.Printf("%s\t", host)
	}
	ips := []string{}
	for _, ip := range IPs {
		ips = append(ips, ip.String())
	}
	fmt.Println(strings.Join(ips, " "))
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
