package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"github.com/golang/glog"
)

// server options
var serverOptions ServerOptions

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

	vendorVmwareNsx := flag.Bool(VENDOR_VMWARE_NSX, false,
		"Include VMware NSX vendor fields.")
	vendorVmwareVds := flag.Bool(VENDOR_VMWARE_VDS, false,
		"Include VMware vSphere Distributed Switch (VDS) vendor fields.")

	exportJsonSyslog := flag.Bool("export-json-to-syslog", false,
		"export flows to syslog server in JSON format")

	exportSyslogAddr := flag.String("export-syslog-host",
		"127.0.0.1",
		"syslog server address for native flow export.")
	exportSyslogPort := flag.Int("export-syslog-port", 514,
		"syslog server port for native flows export.")
	exportSyslogProto := flag.String("export-syslog-proto", "UDP",
		"syslog server proto for native flows export.")
	exportSyslogProgram := flag.String("export-syslog-program",
		"ipfix-forwarder",
		"syslog message program for native flows export.")

	flag.Usage = usage
	flag.Parse()

	displayHeader()

	// display this program version, if requested, and exit.
	if *versionFlag {
		fmt.Println("Git Commit:", GitCommit)
		fmt.Println("Version:", Version)
		if VersionPrerelease != "" {
			fmt.Println("Version PreRelease:", VersionPrerelease)
		}
		os.Exit(2)
	}

	// vendors related options
	var vendors []string
	if *vendorVmwareNsx {
		glog.Infoln("VMWare vendor fields for NSX Netflow will be " +
			"interpreted.")
		vendors = append(vendors, VENDOR_VMWARE_NSX)
	}
	if *vendorVmwareVds {
		glog.Infoln("VMWare vendor fields for vCenter DVS will be " +
			"interpreted.")
		vendors = append(vendors, VENDOR_VMWARE_VDS)
	}

	// syslog server information for native export flows
	var exportSyslogInfo ExportSyslogInfo
	if *exportJsonSyslog {
		exportSyslogInfo = ExportSyslogInfo{
			address: *exportSyslogAddr,
			port:    *exportSyslogPort,
			proto:   strings.ToLower(*exportSyslogProto),
			program: *exportSyslogProgram,
		}
		glog.Infoln("Export to syslog is ON. Destination",
			exportSyslogInfo.proto, "://", exportSyslogInfo.address, ":",
			exportSyslogInfo.port, "program=>", exportSyslogInfo.program)
	} else {
		glog.Infoln("Export to syslog is OFF.")
	}

	// global server options
	serverOptions = ServerOptions{
		address:          *serverAddr,
		port:             *serverPort,
		exportSyslogInfo: exportSyslogInfo,
		vendors:          vendors,
	}

}
