package main

import (
	"log/syslog"
	"github.com/golang/glog"
	"strconv"
	"strings"
)

// sends `msg` string to a syslog server
func sendToSyslog(msg string) {

	syslogInfo := serverOptions.exportSyslogInfo
	syslogService := syslogInfo.address + ":" + strconv.Itoa(syslogInfo.port)

	w, err := syslog.Dial(strings.ToLower(syslogInfo.proto),
		syslogService, syslog.LOG_NOTICE, syslogInfo.program)
	if err != nil {
		glog.Errorln("Could not connect to syslog server: export will fail...")
	} else {

		glog.V(1).Info("Sending JSON message to syslog server:", msg)
		w.Notice(msg)
	}
	defer w.Close()
}
