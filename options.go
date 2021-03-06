package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang/glog"
	"runtime"
)

// server options
var globalServerOptions ServerOptions

func usage() {
	fmt.Println()
	fmt.Fprintln(os.Stderr, "usage: ipfix-forwarder [server-flags] "+
		"[vendor(s)] [syslog-export-info] [logging-properties]")
	fmt.Println()
	flag.PrintDefaults()
	os.Exit(2)
}

// parse command-line flags and return program options
func parseOptions() {

	versionFlag := flag.Bool("version", false, "Version")

	serverAddr := flag.String("server-address", "0.0.0.0",
		"IP the server will be listening to.")
	serverPort := flag.Int("server-port", 2055,
		"Port we will be listening on.")
	serverRcvBuf := flag.Int("server-rcvbuf", 2097152,
		"Size of OS receive buffer associated with the connection.")
	serverSndBuf := flag.Int("server-sndbuf", 2097152,
		"Size of OS transmit buffer associated with the connection.")

	numCPU := flag.Int("num-cpu", runtime.NumCPU(),
		"Number of CPUs to leverage.")

	vendorVmwareNsx := flag.Bool(VendorVmwareNSX, false,
		"Include VMware NSX vendor fields.")
	vendorVmwareVds := flag.Bool(VendorVmwareVDS, false,
		"Include VMware vSphere Distributed Switch (VDS) vendor fields.")
	vendorNokia := flag.Bool(VendorNokia, false,
		"Include Nokia NAT vendor fields.")

	exportJSONSyslog := flag.Bool("export-json-to-syslog", false,
		"export flows to syslog server in JSON format")

	exportSyslogAddr := flag.String("export-syslog-host",
		"127.0.0.1",
		"syslog server address for JSON exports.")
	exportSyslogPort := flag.Int("export-syslog-port", 514,
		"syslog server port forJSON exports.")
	exportSyslogProto := flag.String("export-syslog-proto", "UDP",
		"syslog server proto for JSON exports.")
	exportSyslogProgram := flag.String("export-syslog-program",
		"ipfix-forwarder",
		"syslog message program for JSON exports.")

	flag.Usage = usage
	flag.Parse()

	// display this program version, if requested, and exit.
	if *versionFlag {
		displayVersion()
		os.Exit(2)
	}

	displayHeader()

	// vendors related options
	var vendors []string
	if *vendorVmwareNsx {
		glog.Infoln("VMWare vendor fields for NSX Netflow will be " +
			"interpreted.")
		vendors = append(vendors, VendorVmwareNSX)
	}
	if *vendorVmwareVds {
		glog.Infoln("VMWare vendor fields for vCenter DVS will be " +
			"interpreted.")
		vendors = append(vendors, VendorVmwareVDS)
	}
	if *vendorNokia {
		glog.Infoln("Nokia vendor fields for NAT will be " +
			"interpreted.")
		vendors = append(vendors, VendorNokia)
	}

	// syslog server information for JSON exports.
	var exportSyslogInfo ExportSyslogInfo
	if *exportJSONSyslog {
		exportSyslogInfo = ExportSyslogInfo{
			address: *exportSyslogAddr,
			port:    *exportSyslogPort,
			proto:   strings.ToLower(*exportSyslogProto),
			program: *exportSyslogProgram,
		}
		glog.Infoln("Export to syslog is ON. destination=",
			exportSyslogInfo.proto, "://", exportSyslogInfo.address, ":",
			exportSyslogInfo.port, "program:", exportSyslogInfo.program)
	} else {
		glog.Infoln("Export to syslog is OFF.")
	}

	// global server options
	globalServerOptions = ServerOptions{
		address:          *serverAddr,
		port:             *serverPort,
		rcvbuf:           *serverRcvBuf,
		sndbuf:           *serverSndBuf,
		numCPU:           *numCPU,
		exportSyslogInfo: exportSyslogInfo,
		vendors:          vendors,
	}

}
