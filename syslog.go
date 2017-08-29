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
	syslogWriter     *syslog.Writer
	syslogWriterOnce sync.Once

	syslogErrMsg = "could not send message: syslog connection is nil"
)

// initialize the syslog connection
func getSyslogWriter() (*syslog.Writer, error) {
	var err error
	syslogWriterOnce.Do(func() {
		info := globalServerOptions.exportSyslogInfo
		connStr := info.address + ":" + strconv.Itoa(info.port)
		syslogWriter, err = syslog.Dial(strings.ToLower(info.proto),
			connStr, syslog.LOG_NOTICE, info.program)
		if err == nil {
			defer syslogWriter.Close()
		}
	})
	return syslogWriter, err
}

// is syslog export enabled?
func isSyslogExportEnabled() (bool) {
	return globalServerOptions.exportSyslogInfo != ExportSyslogInfo{}
}

// export message
func exportSyslog(jsonStr string) (error) {
	var errCount uint
	for {
		w, err := getSyslogWriter()
		if err != nil {
			if errCount >= maxRetries {
				return err
			}
			incErrorCountAndSleep(err, &errCount)
			continue
		}
		if w == nil {
			if errCount >= maxRetries {
				return errors.New(syslogErrMsg)
			}

			incErrorCountAndSleep(errors.New(syslogErrMsg), &errCount)
			continue
		}
		glog.V(1).Info(
			"Sending JSON message to syslog server:", jsonStr)
		return w.Notice(jsonStr)
	}
}
