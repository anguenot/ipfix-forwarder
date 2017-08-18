package main

import "github.com/calmh/ipfix"

// ServerOptions contains general `ipfix-forwarder` server options
type ServerOptions struct {
	address          string
	port             int
	vendors          []string
	exportSyslogInfo ExportSyslogInfo
}

// ExportSyslogInfo contains syslog information for native flow exports
type ExportSyslogInfo struct {
	address string
	port    int
	proto   string
	program string
}

// IpfixContext contains an IPFIX session and interpreter
type IpfixContext struct {
	session     *ipfix.Session
	interpreter *ipfix.Interpreter
}
