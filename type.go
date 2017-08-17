package main

// server options
type ServerOptions struct {
	address          string
	port             int
	vendors          []string
	exportSyslogInfo ExportSyslogInfo
}

// syslog information for native flow exports
type ExportSyslogInfo struct {
	address string
	port    int
	proto   string
	program string
}
