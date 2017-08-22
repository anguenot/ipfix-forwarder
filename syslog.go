package main

import (
	"log/syslog"
	"errors"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/glog"
)

var (
	syslogWriter *syslog.Writer
	once         sync.Once
)

// initialize the syslog connection
func getSyslogWriter() *syslog.Writer {
	once.Do(func() {
		info := serverOptions.exportSyslogInfo
		connStr := info.address + ":" + strconv.Itoa(info.port)
		var err error
		syslogWriter, err = syslog.Dial(strings.ToLower(info.proto),
			connStr, syslog.LOG_NOTICE, info.program)
		if err != nil {
			// do not panic here.
			glog.Errorln(err)
		}
		defer syslogWriter.Close()
	})
	return syslogWriter
}

// sends `msg` string to a syslog server
func sendToSyslog(msg string) (error) {
	syslogWriter = getSyslogWriter()
	if syslogWriter != nil {
		glog.V(1).Info("Sending JSON message to syslog server:", msg)
		return syslogWriter.Notice(msg)
	}
	return errors.New("could not send message: syslog connection is nil")
}

// export message
func exportSyslog(jsonStr string) {
	if &serverOptions.exportSyslogInfo == nil {
		return
	}
	if len(jsonStr) > 0 {
		glog.V(2).Infoln("MSG JSON:", jsonStr)
		sendToSyslog(jsonStr)
	} else {
		glog.V(4).Infoln("Empty JSON message: not forwarding.")
	}
}
