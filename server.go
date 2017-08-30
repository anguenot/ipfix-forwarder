package main

import (
	"net"
	"strconv"
	"github.com/golang/glog"
)

// Server holds one (1) UDP connection. Messages are read from the 'incoming'
// channel to be parsed and processed and exporters will read from the
// 'outgoing' channel for exports
type Server struct {
	conn     *net.UDPConn
	incoming chan *Message
	outgoing chan string
	exit     chan interface{}
}

// NewServer returns a new server instance with an active UDP connection.
func NewServer() *Server {
	server := &Server{
		conn:     nil,
		incoming: make(chan *Message),
		outgoing: make(chan string),
		exit:     make(chan interface{}),
	}
	server.Listen()
	for server.conn == nil {
	} // wait for connection
	return server
}

// Listen does wait for incoming UDP messages
func (server *Server) Listen() {

	service := globalServerOptions.address + ":" + strconv.Itoa(globalServerOptions.port)
	udpAddr, _ := net.ResolveUDPAddr("udp", service)

	go func() {

		glog.Infoln("UDP server up and listening on port", string(service))
		glog.Infoln("It can take up to 1 minute for messages to start " +
			"coming in: waiting for IPFIX template sync.")

		var err error
		server.conn, err = net.ListenUDP("udp", udpAddr)
		if err != nil {
			panic(err)
		}
		defer server.conn.Close()

		for {
			select {
			case in := <-server.incoming:
				server.Parse(in)
			case out := <-server.outgoing:
				server.Export(out)
			case <-server.exit:
				return
			}
		}
	}()

}

// read UDP messages and process their IPFIX payload given an IPFIX context
func (server *Server) Read(ipfixContext *IpfixContext, exit chan interface{}) {

	err := error(nil)
	var errCount uint // error count for retry mechanism
	for err == nil && errCount < maxRetries {

		buf := make([]byte, 65507) // maximum UDP payload length
		n, addr, err := server.conn.ReadFrom(buf)
		if err != nil {
			incErrorCountAndSleep(err, &errCount)
			// error will be logged when exiting after 3 errors.
			continue
		}

		glog.V(3).Infoln("Incoming message from UDP client @ ", addr)
		glog.V(3).Infoln("Number of bytes: ", n)

		server.incoming <- NewMessage(ipfixContext, buf, n)

	}

	glog.Errorln("Listener failed 3 times. Killing it!", err)

	exit <- struct{}{}
}

// Parse parses 'msg' and sends JSON representation to 'outgoing' channel.
func (server *Server) Parse(msg *Message) {
	// parse, pre-process and generate a JSON representation.
	go func() {
		server.outgoing <- msg.Parse()
	}()
}

// Start launches one or multiple goroutine listeners.
func (server *Server) Start() {
	glog.Infof("Will be using %d CPU(s).", globalServerOptions.numCPU)
	for cpu := 0; cpu < globalServerOptions.numCPU; cpu++ {
		// use closures with goroutines to ensure we have one (1) IPFIX
		// session and interpreter instances per goroutine
		ipfixContext := initIpfixContext()
		go func(cpu int) {
			glog.Infof("Starting worker #%d ", cpu)
			server.Read(ipfixContext, server.exit)
		}(cpu)
	}
}

// Export sends JSON processed messages to enabled export destinations
func (server *Server) Export(msg string) {
	// syslog export
	if len(msg) > 0 && isSyslogExportEnabled() {
		go exportSyslog(msg)
	}
	// other exports can can here
}

// Message holds an IPFIX context, a buffer containing IPFIX binary payload and
// the size of the buffer.
type Message struct {
	ipfixContext *IpfixContext
	buf          []byte
	n            int
}

// NewMessage returns a *Message
func NewMessage(ipfixContext *IpfixContext, buf []byte, n int) *Message {
	message := &Message{
		ipfixContext: ipfixContext,
		buf:          buf,
		n:            n,
	}
	return message
}

// Parse returns a message in JSON format
func (message *Message) Parse() string {
	return parseIpfix(message.buf, message.n, message.ipfixContext)
}
