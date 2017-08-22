package main

import (
	"net"
	"strconv"

	"github.com/golang/glog"
)

// UDP server
func server() {

	service := serverOptions.address + ":" + strconv.Itoa(serverOptions.port)
	udpAddr, _ := net.ResolveUDPAddr("udp", service)

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	glog.Infoln("UDP server up and listening on port", string(service))
	glog.Infoln("It can take up to 1 minute for messages to start " +
		"coming in: waiting for IPFIX template sync.")

	glog.Infof("Will be using %d CPU(s).", serverOptions.numCPU)

	exit := make(chan struct{})
	for cpu := 0; cpu < serverOptions.numCPU; cpu++ {
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
	for err == nil {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			// error will be logged when exiting
			continue
		}

		glog.V(3).Infoln("Incoming message from UDP client @ ", addr)
		glog.V(3).Infoln("Number of bytes: ", n)

		// parse, pre-process and generate a JSON representation.
		jsonStr := parseIpfix(buf, n, ipfixContext)

		// exports
		go exportSyslog(jsonStr)

	}

	glog.Errorln("A listener died - ", err)

	exit <- struct{}{}

}
