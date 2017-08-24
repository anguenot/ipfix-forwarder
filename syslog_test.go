package main

import (
	"testing"
	"sync"

	"github.com/stretchr/testify/assert"
)

func TestSyslogWriter(t *testing.T) {

	assert.Nil(t, syslogWriter)

	writer, _ := getSyslogWriter()
	keep := writer
	assert.NotNil(t, writer)
	assert.NotNil(t, syslogWriter)
	assert.Equal(t, writer, syslogWriter)

	// force nil
	syslogWriter = nil

	// test init cannot be done twice
	writer, _ = getSyslogWriter()
	assert.Nil(t, writer)
	assert.Nil(t, syslogWriter)
	assert.Equal(t, writer, syslogWriter)

	// re-initialize
	syslogWriterOnce = sync.Once{}
	// force nil
	serverOptions.exportSyslogInfo = ExportSyslogInfo{
		address: "1.2.3.4",
		port:    22,
		proto:   "UDP",
		program: "ipfix-forwarder",
	}
	w, err := getSyslogWriter()
	// always returns a writer even with bad data.
	assert.NotNil(t, w)
	assert.Nil(t, err)

	// put back
	syslogWriter = keep

}

func TestExportSyslogError(t *testing.T) {
	// force nil
	syslogWriter = nil
	err := exportSyslog("testing")
	assert.Error(t, err, syslogErrMsg)
}

func TestExportSyslogEmptyError(t *testing.T) {
	// force nil
	syslogWriter = nil
	err := exportSyslog("")
	assert.Error(t, err, syslogErrMsg)
}

func TestExportSyslogWriterError(t *testing.T) {
	serverOptions.exportSyslogInfo = ExportSyslogInfo{
		address: "1.2.3.4",
		port:    2055,
		proto:   "UDP",
		program: "testing",
	}
	err := exportSyslog("whatever")
	assert.Error(t, err, syslogErrMsg)

	serverOptions.exportSyslogInfo = ExportSyslogInfo{}
}

func TestSyslogExportDisabled(t *testing.T) {

	assert.False(t, isSyslogExportEnabled())

	serverOptions = ServerOptions{}
	assert.False(t, isSyslogExportEnabled())

	serverOptions.exportSyslogInfo = ExportSyslogInfo{}
	assert.False(t, isSyslogExportEnabled())

	serverOptions.exportSyslogInfo = ExportSyslogInfo{
		address: "",
		port:    0,
		proto:   "",
		program: "",
	}
	// still null cause empty struct
	assert.False(t, isSyslogExportEnabled())

	serverOptions.exportSyslogInfo = ExportSyslogInfo{
		address: "",
		port:    2055,
		proto:   "",
		program: "",
	}
	// still null cause empty struct
	assert.True(t, isSyslogExportEnabled())

	// put it back to disabled for other tests
	serverOptions.exportSyslogInfo = ExportSyslogInfo{}

}

func TestSyslogExportEnabled(t *testing.T) {
	assert.False(t, isSyslogExportEnabled())
}

func TestExportSyslog(t *testing.T) {

	// re-initialize
	syslogWriterOnce = sync.Once{}

	// force nil
	serverOptions.exportSyslogInfo = ExportSyslogInfo{
		address: "localhost",
		port:    2055,
		proto:   "UDP",
		program: "ipfix-forwarder",
	}
	err := exportSyslog("{}")
	assert.Nil(t, err)

	syslogWriter = nil
}
