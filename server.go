package main

import (
	"net"
	"log"
	"github.com/golang/glog"
	"strconv"
	"runtime"
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

	exit := make(chan struct{})
	for cpu := 0; cpu < runtime.NumCPU(); cpu++ {
		go readUDP(conn, exit)
	}
	<-exit

}

// read UDP message
func readUDP(conn *net.UDPConn, exit chan struct{}) {

	buf := make([]byte, 65507) // maximum UDP payload length

	err := error(nil)
	for err == nil {
		n, addr, err := conn.ReadFrom(buf)
		glog.V(3).Infoln("Incoming message from UDP client @ ", addr)
		glog.V(3).Infoln("Number of bytes: ", n)
		if err != nil {
			log.Fatal(err)
		}
		jsonStr := parseIpfix(buf, n)
		if &serverOptions.exportSyslogInfo != nil {
			go exportSyslog(jsonStr)
		}
	}

	glog.Errorln("A listener died - ", err)
	exit <- struct{}{}
}

// parse IPFIX messages and returns a JSON string representation
func parseIpfix(buf []byte, n int) (string) {
	msgMap := parseIpfixMessage(buf, n)
	var jsonStr string
	if len(msgMap) > 0 {
		jsonStr = mapToJSON(msgMap)
	} else {
		glog.V(3).Infoln("Empty message: waiting for schema?")
	}
	return jsonStr
}

// export message
func exportSyslog(jsonStr string) {
	if len(jsonStr) > 0 {
		glog.V(2).Infoln("MSG JSON:", jsonStr)
		sendToSyslog(jsonStr)
	} else {
		glog.V(4).Infoln("Empty JSON message: not forwarding.")
	}
}
