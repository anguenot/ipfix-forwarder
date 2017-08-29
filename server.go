package main

import (
	"net"
	"strconv"
	"github.com/golang/glog"
)

// UDP server
func server() {

	service := globalServerOptions.address + ":" + strconv.Itoa(globalServerOptions.port)
	udpAddr, _ := net.ResolveUDPAddr("udp", service)

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	glog.Infoln("UDP server up and listening on port", string(service))
	glog.Infoln("It can take up to 1 minute for messages to start " +
		"coming in: waiting for IPFIX template sync.")

	glog.Infof("Will be using %d CPU(s).", globalServerOptions.numCPU)

	exit := make(chan struct{})
	for cpu := 0; cpu < globalServerOptions.numCPU; cpu++ {
		// use closures with goroutines to ensure we have one (1) IPFIX
		// session and interpreter instances per goroutine
		ipfixContext := initIpfixContext()
		go func(cpu int) {
			glog.Infof("Starting listener #%d ", cpu)
			readUDP(conn, ipfixContext, exit)
		}(cpu)
	}
	<-exit

}

// read UDP message
func readUDP(conn *net.UDPConn, ipfixContext *IpfixContext,

	exit chan struct{}) {

	buf := make([]byte, 65507) // maximum UDP payload length

	err := error(nil)
	var errCount uint // error count for retry mechanism
	for err == nil && errCount < maxRetries {

		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			incErrorCountAndSleep(err, &errCount)
			// error will be logged when exiting after 3 errors.
			continue
		}

		glog.V(3).Infoln("Incoming message from UDP client @ ", addr)
		glog.V(3).Infoln("Number of bytes: ", n)

		// parse, pre-process and generate a JSON representation.
		jsonStr := parseIpfix(buf, n, ipfixContext)

		// exports
		if len(jsonStr) > 0 && isSyslogExportEnabled() {
			go exportSyslog(jsonStr)
		}

	}

	glog.Errorln("Listener failed 3 times. Killing it!", err)

	exit <- struct{}{}

}
